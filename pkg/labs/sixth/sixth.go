package sixth

import (
	"fmt"
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
