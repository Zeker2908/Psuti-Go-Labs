package seventh

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем метод, URL и начало выполнения
		log.Printf("Начало запроса: %s %s", r.Method, r.URL.Path)

		// Вызываем следующий обработчик
		next.ServeHTTP(w, r)

		// Логируем продолжительность выполнения запроса
		log.Printf("Запрос %s %s выполнен за %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// Запуск HTTP-сервера
func HttpServer() {
	mux := http.NewServeMux()

	// Регистрация маршрутов
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)

	// Оборачиваем маршруты в middleware
	loggedMux := loggingMiddleware(mux)

	// Запуск HTTP-сервера
	fmt.Println("HTTP-сервер запущен на порту 8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal("Ошибка при запуске HTTP-сервера:", err)
	}
}

// Обработчик для GET-запроса на /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	// Возвращаем приветственное сообщение
	fmt.Fprintln(w, "Привет! Это ответ на GET-запрос.")
}

// Обработчик для POST-запроса на /data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Парсинг данных JSON
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	// Вывод данных в консоль
	fmt.Println("Полученные данные:", data)

	// Отправляем подтверждение клиенту
	fmt.Fprintln(w, "Данные получены успешно")
}

func TcpServer(wg *sync.WaitGroup) {
	// Указываем порт для прослушивания
	port := ":4545"

	// Запуск сервера
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен и слушает порт", port)

	// Канал для получения системных сигналов (например, для завершения)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Канал для закрытия сервера
	stop := make(chan struct{})

	// Горутин для прослушивания соединений
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				select {
				case <-stop: // Проверяем, нужно ли остановить сервер
					return
				default:
					fmt.Println("Ошибка при принятии соединения:", err)
					continue
				}
			}

			wg.Add(1) // Добавляем горутину в группу ожидания
			go handleConnection(conn, wg)
		}
	}()

	// Ожидание сигнала завершения работы
	<-shutdown
	fmt.Println("Завершение работы сервера...")

	// Останавливаем сервер
	close(stop)

	// Ожидание завершения всех горутин
	wg.Wait()

	fmt.Println("Сервер завершил работу")
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done() // Сообщаем о завершении работы горутины

	// Чтение сообщения от клиента
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении сообщения:", err)
		return
	}

	// Выводим сообщение на экран
	fmt.Println("Получено сообщение:", message)

	// Отправляем ответ клиенту
	_, err = conn.Write([]byte("Сообщение получено\n"))
	if err != nil {
		fmt.Println("Ошибка при отправке ответа:", err)
		return
	}
}
