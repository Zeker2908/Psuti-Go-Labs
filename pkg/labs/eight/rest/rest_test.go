package rest

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(errorHandler())
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	return router
}

func setupTestDB() {
	config := &Config{
		Database: struct {
			User     string `yaml:"user"`
			Dbname   string `yaml:"dbname"`
			Sslmode  string `yaml:"sslmode"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
		}{
			User:     "leonidas",
			Dbname:   "test_for_test",
			Sslmode:  "disable",
			Password: "leo29082004",
			Host:     "localhost",
			Port:     "5432",
		},
	}

	db = connectToDB(config)
	createTable(db)
	validate = validator.New()

}

func TestCreateUser(t *testing.T) {
	// Set up test database
	setupTestDB()
	defer db.Close()

	router := setupRouter()

	user := User{Name: "John Doe", Email: "john@example.com"}
	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdUser User
	json.Unmarshal(w.Body.Bytes(), &createdUser)

	// Ensure user is created with valid ID
	assert.NotZero(t, createdUser.ID)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestGetUser(t *testing.T) {
	// Set up test database
	setupTestDB()
	defer db.Close()

	router := setupRouter()

	// Create a user to get
	user := User{Name: "Jane Doe", Email: "jane@example.com"}
	db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)

	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(user.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedUser User
	json.Unmarshal(w.Body.Bytes(), &retrievedUser)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, user.Name, retrievedUser.Name)
	assert.Equal(t, user.Email, retrievedUser.Email)
}
func TestUpdateUser(t *testing.T) {
	setupTestDB()
	defer db.Close()

	router := setupRouter()

	// Создаём пользователя, которого будем обновлять
	originalUser := User{Name: "Alice", Email: "alice@example.com"}
	db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", originalUser.Name, originalUser.Email).Scan(&originalUser.ID)

	// Обновляемые данные
	updatedUser := User{Name: "Alice Smith", Email: "alice.smith@example.com"}
	updatedUserJSON, _ := json.Marshal(updatedUser)

	// Отправляем запрос на обновление
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(originalUser.ID), bytes.NewBuffer(updatedUserJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что пользователь был обновлён в базе данных
	var retrievedUser User
	db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", originalUser.ID).Scan(&retrievedUser.ID, &retrievedUser.Name, &retrievedUser.Email)

	// Проверяем обновлённые данные
	assert.Equal(t, updatedUser.Name, retrievedUser.Name)
	assert.Equal(t, updatedUser.Email, retrievedUser.Email)
}

func TestDeleteUser(t *testing.T) {
	// Set up test database
	setupTestDB()
	defer db.Close()

	router := setupRouter()

	// Create a user to delete
	user := User{Name: "Mark Smith", Email: "mark@example.com"}
	db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(user.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Verify the user has been deleted
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", user.ID).Scan(&count)
	assert.Equal(t, 0, count)
}
