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
