package main

import (
	"PsutiGoLabs/pkg/labs/first"
	"PsutiGoLabs/pkg/labs/second"
	"fmt"
)

func main() {
	selectLabsAndTasks()
}

func selectLabsAndTasks() {
	var labNumber, taskNumber int
	fmt.Print("Введите номер лабораторной работы: ")
	fmt.Scanln(&labNumber)

	switch labNumber {
	case 1:
		fmt.Print("Введите номер задания (1-6): ")
		fmt.Scanln(&taskNumber)

		switch taskNumber {
		case 1:
			first.WhatTimeIsIt()
		case 2, 3:
			first.PrintOfNumbers()
		case 4:
			var a, b int
			fmt.Print("Введите два целых числа: ")
			fmt.Scanln(&a, &b)
			first.CalculateInt(a, b)
		case 5:
			var a, b float64
			fmt.Print("Введите два числа с плавающей точкой: ")
			fmt.Scanln(&a, &b)
			first.CalculateFloat(a, b)
		case 6:
			var a, b, c int
			fmt.Print("Введите три числа с плавающей точкой: ")
			fmt.Scanln(&a, &b, &c)
			first.Average(a, b, c)
		default:
			fmt.Println("Неверный номер задания")
		}
	case 2:
		fmt.Print("Введите номер задания (1-6): ")
		fmt.Scanln(&taskNumber)

		switch taskNumber {
		case 1:
			var a int
			fmt.Print("Введите число: ")
			fmt.Scanln(&a)
			fmt.Println("Проверка на четность: ", second.Parity(a))
		case 2:
			var a int
			fmt.Print("Введите число: ")
			fmt.Scanln(&a)
			fmt.Println(second.CheckNumberSign(a))
		case 3:
			second.PrintNumbers()
		case 4:
			var str string
			fmt.Print("Введите строку: ")
			fmt.Scanln(&str)
			fmt.Println(second.StringLength(str))
		case 5:
			var width, height int
			fmt.Print("Введите ширину и высоту прямоугольника: ")
			fmt.Scanln(&width, &height)
			rectangle := second.Rectangle{Width: width, Height: height}
			fmt.Printf("Площадь вашего прямоугольника равна = %d\n", second.Square(rectangle))
		case 6:
			var a, b int
			fmt.Print("Введите два целых числа: ")
			fmt.Scanln(&a, &b)
			fmt.Printf("Среднее арифмитическое = %f\n", second.Average(a, b))

		default:
			fmt.Println("Неверный номер задания")
		}

	default:
		fmt.Println("Неверный номер лабораторной работы")
	}
}
