package sixth

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1) Создание и запуск горутин:
// Функция для расчёта факториала
func FactorialSync(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
		time.Sleep(100 * time.Millisecond) // Имитация задержки
	}
	fmt.Printf("Факториал %d равен %d\n", n, result)
}

// Функция для генерации случайных чисел
func GenerateRandomNumbersSync(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Случайное число: %d\n", num)
		time.Sleep(150 * time.Millisecond) // Имитация задержки
	}
}

// Функция для вычисления суммы числового ряда
func SumSeriesSync(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(120 * time.Millisecond) // Имитация задержки
	}
	fmt.Printf("Сумма числового ряда до %d равна %d\n", n, sum)
}

// 2) Использование каналов для передачи данных
func FibonacciInput(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fib := []int{0, 1}
	for i := 0; i < 10; i++ {
		fib = append(fib, fib[i]+fib[i+1])
	}
	ch <- fib
	fmt.Println("Первая горутина закончила работу")
}
func FibonacciOut(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := <-ch
	fmt.Println(result)
	fmt.Println("Вторая горутина закончила работу")
}

// 3) Применение select для управления каналами
func GenerateNumbers(numbers chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) // генерируем случайное число от 0 до 99
		numbers <- num
		time.Sleep(500 * time.Millisecond) // небольшая задержка для демонстрации
	}
}

func CheckEvenOdd(num int, results chan<- string) {
	if num%2 == 0 {
		results <- fmt.Sprintf("%d is even", num)
	} else {
		results <- fmt.Sprintf("%d is odd", num)
	}
}

// 4.Синхронизация с помощью мьютексов:
var Counter int

func Increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mutex.Lock() // блокируем доступ к переменной
		Counter += 1
		mutex.Unlock() // разблокируем доступ
	}
}

// 5. Разработка многопоточного калькулятора:

// Структура запроса калькулятора
type CalcRequest struct {
	Operation string       // Операция: "+", "-", "*", "/"
	A, B      float64      // Операнды
	Result    chan float64 // Канал для отправки результата
}

// Функция калькулятора
func Calculator(wg *sync.WaitGroup, requests <-chan CalcRequest) {
	defer wg.Done()

	for req := range requests {
		var result float64
		switch req.Operation {
		case "+":
			result = req.A + req.B
		case "-":
			result = req.A - req.B
		case "*":
			result = req.A * req.B
		case "/":
			if req.B != 0 {
				result = req.A / req.B
			} else {
				log.Println("Ошибка: деление на ноль")
				result = 0
			}
		default:
			log.Println("Ошибка: неподдерживаемая операция")
			result = 0
		}
		req.Result <- result // Отправка результата через канал
	}
}

//6) Создание пула воркеров

// Структура задачи
type Task struct {
	ID      int
	Payload int // Данные задачи (например, число для обработки)
}

// Функция воркера
func Worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		// Обработка задачи (например, удвоение числа)
		result := fmt.Sprintf("Worker %d обработал задачу %d: %d * 2 = %d", id, task.ID, task.Payload, task.Payload*2)
		results <- result // Отправляем результат
	}
}
