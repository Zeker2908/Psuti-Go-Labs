package main

import (
	"PsutiGoLabs/pkg/labs/fifth"
	"PsutiGoLabs/pkg/labs/first"
	"PsutiGoLabs/pkg/labs/fourth"
	"PsutiGoLabs/pkg/labs/second"
	"PsutiGoLabs/pkg/labs/sixth"
	"PsutiGoLabs/pkg/labs/third"
	"PsutiGoLabs/pkg/labs/third/mathutils"
	"PsutiGoLabs/pkg/labs/third/stringutils"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	selectLabsAndTasks()
}

func selectLabsAndTasks() {
	labNumber := scanInt("Введите номер лабораторной работы: ")
	taskNumber := scanInt("Введите номер задания (1-6): ")

	switch labNumber {
	case 1:
		handleFirstLab(taskNumber)
	case 2:
		handleSecondLab(taskNumber)
	case 3:
		handleThirdLab(taskNumber)
	case 4:
		handleFourthLab(taskNumber)
	case 5:
		handleFifthLab(taskNumber)
	case 6:
		handleSixthLab(taskNumber)
	default:
		fmt.Println("Неверный номер лабораторной работы")
	}
}

func handleFirstLab(taskNumber int) {
	switch taskNumber {
	case 1:
		first.WhatTimeIsIt()
	case 2, 3:
		first.PrintOfNumbers()
	case 4:
		a, b := scanTwoInts("Введите два целых числа: ")
		first.CalculateInt(a, b)
	case 5:
		a, b := scanTwoFloats("Введите два числа с плавающей точкой: ")
		first.CalculateFloat(a, b)
	case 6:
		a, b, c := scanThreeInts("Введите три целых числа: ")
		first.Average(a, b, c)
	default:
		fmt.Println("Неверный номер задания")
	}
}

func handleSecondLab(taskNumber int) {
	switch taskNumber {
	case 1:
		a := scanInt("Введите число: ")
		fmt.Println("Проверка на четность: ", second.Parity(a))
	case 2:
		a := scanInt("Введите число: ")
		fmt.Println(second.CheckNumberSign(a))
	case 3:
		second.PrintNumbers()
	case 4:
		str := scanString("Введите строку: ")
		fmt.Println(second.StringLength(str))
	case 5:
		width, height := scanTwoInts("Введите ширину и высоту прямоугольника: ")
		rectangle := second.NewRectangle(width, height)
		fmt.Printf("Площадь вашего прямоугольника равна = %d\n", rectangle.Area())
	case 6:
		a, b := scanTwoInts("Введите два целых числа: ")
		fmt.Printf("Среднее арифмитическое = %f\n", second.Average(a, b))
	default:
		fmt.Println("Неверный номер задания")
	}
}

func handleThirdLab(taskNumber int) {
	switch taskNumber {
	case 1, 2:
		num := scanInt("Введите число для вычисления факториала: ")
		fmt.Println(mathutils.Factorial(num))
	case 3:
		str := scanString("Введите строку, которую хотите перевернуть: ")
		fmt.Println(stringutils.Reverse(str))
	case 4:
		third.CreateAndPrintArray()
	case 5:
		third.MakeSlice()
	case 6:
		fmt.Println(third.FindLongestString("str", "hello world", "Nice", "  "))
	default:
		fmt.Println("Неверный номер задания")
	}
}

func handleFourthLab(taskNumber int) {
	switch taskNumber {
	case 1:
		people := createPeopleMap()
		fourth.PrintMap(people)
	case 2:
		people := createPeopleMap()
		fmt.Println(fourth.AvgAgePeopleMap(people))
	case 3:
		people := createPeopleMap()
		fourth.PrintMap(people)
		fourth.DeleteElement(people, "Alice")
		fmt.Println()
		fourth.PrintMap(people)
	case 4:
		str := scanString("Введите строку: ")
		fmt.Println(strings.ToUpper(str))
	case 5:
		numbers := FillingArr()
		fmt.Printf("Сумма введенных чисел: %d\n", fourth.Sum(numbers...))
	case 6:
		numbers := fourth.IntReverse(FillingArr())
		fmt.Println(numbers)
	default:
		fmt.Println("Неверный номер задания")
	}
}

func handleFifthLab(taskNumber int) {
	switch taskNumber {
	case 1:
		person := fifth.NewPerson("Alex", 15)
		fmt.Println(person)
	case 2:
		person := fifth.NewPerson("Bobi", 55)
		person.Birthday()
		fmt.Println(person)
	case 4, 5:
		rect := second.NewRectangle(3, 6)
		circ := fifth.NewCircle(6)
		shapes := []fifth.Shape{rect, circ}
		fifth.PrintArea(shapes)
	case 6:
		book := fifth.Book{Title: "Face", Author: "Ivan Face", Year: 1997}
		fmt.Println(book)
	default:
		fmt.Println("Неверный номер задания")
	}
}

func handleSixthLab(taskNumber int) {
	switch taskNumber {
	case 1:
		var wg sync.WaitGroup

		// Увеличиваем счетчик WaitGroup на 3
		wg.Add(3)

		// Запуск трех функций параллельно
		go sixth.FactorialSync(5, &wg)
		go sixth.GenerateRandomNumbersSync(5, &wg)
		go sixth.SumSeriesSync(10, &wg)

		// Ожидаем завершения всех задач
		wg.Wait()

		fmt.Println("Все задачи завершены")
	case 2:
		var wg sync.WaitGroup
		ch := make(chan []int)
		wg.Add(2)
		go sixth.FibonacciInput(ch, &wg)
		go sixth.FibonacciOut(ch, &wg)
		wg.Wait()
		close(ch)
	case 3:
		numbers := make(chan int)
		results := make(chan string)

		// Запуск горутины для генерации чисел
		go sixth.GenerateNumbers(numbers)

		// Использование select для управления каналами
		for {
			select {
			case num := <-numbers:
				fmt.Printf("Генерируемое значение: %d\n", num)
				// Проверяем чётность сгенерированного числа в отдельной горутине
				go sixth.CheckEvenOdd(num, results)
			case result := <-results:
				fmt.Println(result)
			}
		}
	case 4:
		var wg sync.WaitGroup
		var mutex sync.Mutex
		numGoroutines := scanInt("Введите количество горутин ")

		// Запускаем несколько горутин
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go sixth.Increment(&wg, &mutex)
		}

		wg.Wait()

		fmt.Printf("Финальное число: %d\n", sixth.Counter)
	case 5:
		requests := make(chan sixth.CalcRequest)
		var wg sync.WaitGroup

		// Запуск нескольких горутин-калькуляторов
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go sixth.Calculator(&wg, requests)
		}

		// Пример клиентских запросов
		go func() {
			for _, req := range []sixth.CalcRequest{
				{"+", 10, 5, make(chan float64)},
				{"-", 20, 4, make(chan float64)},
				{"*", 6, 3, make(chan float64)},
				{"/", 15, 3, make(chan float64)},
				{"/", 15, 0, make(chan float64)}, // Деление на 0
			} {
				requests <- req
				fmt.Printf("Операция: %s, %f %s %f = %f\n", req.Operation, req.A, req.Operation, req.B, <-req.Result)
			}
			close(requests) // Закрытие канала запросов после отправки всех
		}()

		wg.Wait() // Ожидание завершения всех горутин
	case 6:
		// Запрос количества воркеров у пользователя
		numWorkers := scanInt("Введите количество воркеров:")

		// Канал для задач и результатов
		tasks := make(chan sixth.Task, 10)
		results := make(chan string, 10)

		// WaitGroup для ожидания завершения всех воркеров
		var wg sync.WaitGroup

		// Запуск воркеров
		for i := 1; i <= numWorkers; i++ {
			wg.Add(1)
			go sixth.Worker(i, tasks, results, &wg)
		}

		// Генерация задач
		go func() {
			for i := 1; i <= 10; i++ {
				tasks <- sixth.Task{ID: i, Payload: i * 10} // Пример задачи
			}
			close(tasks) // Закрываем канал после отправки всех задач
		}()

		// Ожидание завершения всех воркеров
		go func() {
			wg.Wait()
			close(results) // Закрываем канал результатов после завершения воркеров
		}()

		// Открываем файл для записи результатов
		file, err := os.Create("worker_results.txt")
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		// Получение и вывод результатов в файл и консоль
		for result := range results {
			fmt.Println(result) // Вывод в консоль
			_, err := file.WriteString(result + "\n")
			if err != nil {
				return
			} // Запись в файл
		}

		fmt.Println("Все задачи завершены, результат сохранен в worker_results.txt.")

	default:
		fmt.Println("Неверный номер задания")
	}

}

func FillingArr() []int {
	var numbers []int
	fmt.Println("Введите числа (для завершения ввода введите любой символ):")
	for {
		var number int
		if _, err := fmt.Scan(&number); err != nil {
			break
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func scanInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return 0
	}
	return input
}

func scanTwoInts(prompt string) (int, int) {
	var a, b int
	fmt.Print(prompt)
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return 0, 0
	}
	return a, b
}

func scanTwoFloats(prompt string) (float64, float64) {
	var a, b float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return 0, 0
	}
	return a, b
}

func scanThreeInts(prompt string) (int, int, int) {
	var a, b, c int
	fmt.Print(prompt)
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return 0, 0, 0
	}
	return a, b, c
}

func scanString(prompt string) string {
	var str string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return ""
	}
	return str
}

func createPeopleMap() map[string]int {
	names := []string{"Alice", "Bob", "Charlie"}
	ages := []int{50, 25, 35}
	return fourth.PeopleMap(names, ages)
}
