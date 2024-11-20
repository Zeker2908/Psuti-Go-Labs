package rest

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"` // Роль пользователя: "admin" или "user"
}

var jwtKey = []byte("secret_key") // Этот ключ используется для подписи токенов

// Структура для создания токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Создание JWT токена
func generateJWT(user User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа
	claims := &Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Используем Unix() для получения времени в формате Unix
			Issuer:    "N.I.G.G.A",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Валидация JWT токена
func validateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
func registerUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Назначение роли по умолчанию
	if user.Role == "" {
		user.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хэширования пароля"})
		return
	}
	user.Password = string(hashedPassword)

	err = db.QueryRow("INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Логин пользователя и получение токена
func loginUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверить данные пользователя (например, по email)
	var dbUser User
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email=$1", user.Email).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Email, &dbUser.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	// Генерировать JWT токен
	token, err := generateJWT(dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Middleware для проверки токена
func authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
			c.Abort()
			return
		}

		// Убираем "Bearer " из строки токена
		tokenString = tokenString[len("Bearer "):]

		// Валидация токена
		claims, err := validateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Добавляем данные из токена в контекст
		c.Set("username", claims.Username)
		c.Next()
	}
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
            email VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(50) NOT NULL DEFAULT 'user'
        );
    `
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем администратора, если его нет
	adminExistsQuery := `SELECT COUNT(*) FROM users WHERE email = 'admin@admin.com'`
	var adminCount int
	err = db.QueryRow(adminExistsQuery).Scan(&adminCount)
	if err != nil {
		log.Fatal(err)
	}

	if adminCount == 0 {
		// Хэшируем пароль администратора
		adminPassword := "adminpass"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Ошибка хэширования пароля администратора:", err)
		}

		_, err = db.Exec(
			`INSERT INTO users (name, email, password, role) VALUES ('Admin', 'admin@admin.com', $1, 'admin')`,
			hashedPassword,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Admin user created: email=admin@admin.com, password=adminpass")
	}
}

func authorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
			c.Abort()
			return
		}

		// Убираем "Bearer " из строки токена
		tokenString = tokenString[len("Bearer "):]

		// Валидация токена
		claims, err := validateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Проверяем роль пользователя
		var role string
		err = db.QueryRow("SELECT role FROM users WHERE name = $1", claims.Username).Scan(&role)
		if err != nil || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Получить всех пользователей с пагинацией и фильтрацией
func getUsers(c *gin.Context) {
	// Проверка токена для этого маршрута
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
		return
	}

	// Убираем "Bearer " из строки токена
	tokenString = tokenString[len("Bearer "):]

	// Валидация токена
	_, err := validateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	// Пагинация и фильтрация пользователей
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
		query = "SELECT id, name, email, password FROM users WHERE name ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3"
		countQuery = "SELECT COUNT(*) FROM users WHERE name ILIKE $1"
		args = append(args, "%"+name+"%", limitInt, offset)
	} else {
		query = "SELECT id, name, email, password FROM users ORDER BY id LIMIT $1 OFFSET $2"
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
		// Включаем поле Password в выборку
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	// Отправка ответа с дополнительной информацией о пагинации
	c.JSON(http.StatusOK, gin.H{
		"limit": limitInt,
		"page":  pageInt,
		"total": count,
		"users": users,
	})
}

// Получить одного пользователя по ID
func getUserByID(c *gin.Context) {
	// Проверка токена для этого маршрута
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
		return
	}

	// Убираем "Bearer " из строки токена
	tokenString = tokenString[len("Bearer "):]

	// Валидация токена
	_, err := validateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	id := c.Param("id")
	var user User
	err = db.QueryRow("SELECT id, name, email, password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Обновить пользователя
func updateUser(c *gin.Context) {
	// Проверка токена для этого маршрута
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
		return
	}

	// Убираем "Bearer " из строки токена
	tokenString = tokenString[len("Bearer "):]

	// Валидация токена
	_, err := validateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	id := c.Param("id")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4",
		user.Name, user.Email, user.Password, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
			return
		}
		user.Password = hashedPassword
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь обновлен"})
}

// Удалить пользователя
func deleteUser(c *gin.Context) {
	// Проверка токена для этого маршрута
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
		return
	}

	// Убираем "Bearer " из строки токена
	tokenString = tokenString[len("Bearer "):]

	// Валидация токена
	_, err := validateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		return
	}

	id := c.Param("id")
	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удален"})
}

// Хэширование пароля
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Проверка пароля
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Основная функция
func Start() {
	// Загрузка конфигурации
	config, err := loadConfig("config/local.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к базе данных
	db = connectToDB(config)

	// Создание таблицы пользователей
	createTable(db)

	// Создание экземпляра Gin
	r := gin.Default()

	// Маршрут для регистрации пользователя
	r.POST("/register", registerUser)

	// Маршрут для входа и получения токена
	r.POST("/login", loginUser)

	protected := r.Group("/", authenticate())
	{
		// Получение пользователя по ID
		protected.GET("/users/:id", getUserByID)

		// Создание пользователя (доступно всем авторизованным)
		protected.POST("/users", registerUser)

		// Обновление и удаление доступны только администратору
		protected.PUT("/users/:id", authorizeAdmin(), updateUser)
		protected.DELETE("/users/:id", authorizeAdmin(), deleteUser)
		protected.GET("/users", getUsers)

	}

	// Запуск сервера
	r.Run(":8080")
}
