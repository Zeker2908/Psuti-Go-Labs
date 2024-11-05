package rest

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

var db *sql.DB
var validate *validator.Validate

// Загрузить конфигурацию из YAML файла
func loadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// Подключиться к базе данных
func connectToDB(config *Config) *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s",
		config.Database.User, config.Database.Dbname, config.Database.Sslmode,
		config.Database.Password, config.Database.Host, config.Database.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Создать таблицу пользователей, если она не существует
func createTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL
        );
    `
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
func createTableLite(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

// Получить всех пользователей с пагинацией и фильтрацией
func getUsers(c *gin.Context) {
	// Значения по умолчанию для пагинации
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	name := c.Query("name")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}

	offset := (pageInt - 1) * limitInt

	var query string
	var countQuery string
	var args []interface{}

	if name != "" {
		query = "SELECT id, name, email FROM users WHERE name ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3"
		countQuery = "SELECT COUNT(*) FROM users WHERE name ILIKE $1"
		args = append(args, "%"+name+"%", limitInt, offset)
	} else {
		query = "SELECT id, name, email FROM users ORDER BY id LIMIT $1 OFFSET $2"
		countQuery = "SELECT COUNT(*) FROM users"
		args = append(args, limitInt, offset)
	}

	var count int
	if name != "" {
		err = db.QueryRow(countQuery, "%"+name+"%").Scan(&count)
	} else {
		err = db.QueryRow(countQuery).Scan(&count)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": count,
		"page":  pageInt,
		"limit": limitInt,
	})
}

// Получить пользователя по ID
func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	var user User
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Создать нового пользователя
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверить входные данные пользователя
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Обновить существующего пользователя
func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверить входные данные пользователя
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновить пользователя в базе данных
	_, err = db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удален"})
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0]
			var status int

			if err.Type == gin.ErrorTypePrivate {
				status = http.StatusInternalServerError
			} else {
				status = http.StatusBadRequest
			}

			c.JSON(status, gin.H{"error": err.Error()})
			c.Abort()
		}
	}
}

func Start() {
	config, err := loadConfig("config/local.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db = connectToDB(config)
	defer db.Close()

	createTable(db)

	validate = validator.New()

	router := gin.Default()
	router.Use(errorHandler())

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}
