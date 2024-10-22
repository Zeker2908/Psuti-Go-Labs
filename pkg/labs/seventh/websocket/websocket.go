package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// Мьютекс для синхронизации доступа к клиентам
var mu sync.Mutex

// Мапа для хранения подключённых клиентов
var clients = make(map[*websocket.Conn]bool)

// Канал для передачи сообщений между клиентами
var broadcast = make(chan Message)

// Настройка веб-сокетов
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Структура для сообщений
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func WebSocketServer() {
	// Настройка маршрутов
	http.HandleFunc("/ws", handleConnections)

	// Запуск горутины для обработки сообщений
	go handleMessages()

	// Запуск HTTP-сервера
	fmt.Println("Сервер запущен на :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}

}

// Функция для обработки новых подключений
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Обновление соединения до веб-сокет
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	// Добавляем нового клиента
	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	// Чтение сообщений от клиента
	for {
		var msg Message
		// Чтение сообщения JSON
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v", err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
		// Отправка сообщения в канал broadcast
		broadcast <- msg
	}
}

// Функция для рассылки сообщений всем клиентам
func handleMessages() {
	for {
		// Получение сообщения из канала
		msg := <-broadcast
		// Отправка сообщения всем подключённым клиентам
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Ошибка отправки сообщения клиенту: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
