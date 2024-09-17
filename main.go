package main

import (
	"PsutiGoLabs/pkg/labs/first"
	"PsutiGoLabs/pkg/labs/second"
	"PsutiGoLabs/pkg/labs/third"
	"PsutiGoLabs/pkg/labs/third/mathutils"
	"PsutiGoLabs/pkg/labs/third/stringutils"
	"fmt"
)

func main() {
	selectLabsAndTasks()
}

func selectLabsAndTasks() {
	var labNumber, taskNumber int
	fmt.Print("Введите номер лабораторной работы: ")
	_, err := fmt.Scanln(&labNumber)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	switch labNumber {
	case 1:
		fmt.Print("Введите номер задания (1-6): ")
		_, err := fmt.Scanln(&taskNumber)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}

		switch taskNumber {
		case 1:
			first.WhatTimeIsIt()
		case 2, 3:
			first.PrintOfNumbers()
		case 4:
			var a, b int
			fmt.Print("Введите два целых числа: ")
			_, err := fmt.Scanln(&a, &b)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			first.CalculateInt(a, b)
		case 5:
			var a, b float64
			fmt.Print("Введите два числа с плавающей точкой: ")
			_, err := fmt.Scanln(&a, &b)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			first.CalculateFloat(a, b)
		case 6:
			var a, b, c int
			fmt.Print("Введите три числа с плавающей точкой: ")
			_, err := fmt.Scanln(&a, &b, &c)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			first.Average(a, b, c)
		default:
			fmt.Println("Неверный номер задания")
		}
	case 2:
		fmt.Print("Введите номер задания (1-6): ")
		_, err := fmt.Scanln(&taskNumber)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}

		switch taskNumber {
		case 1:
			var a int
			fmt.Print("Введите число: ")
			_, err := fmt.Scanln(&a)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			fmt.Println("Проверка на четность: ", second.Parity(a))
		case 2:
			var a int
			fmt.Print("Введите число: ")
			_, err := fmt.Scanln(&a)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			fmt.Println(second.CheckNumberSign(a))
		case 3:
			second.PrintNumbers()
		case 4:
			var str string
			fmt.Print("Введите строку: ")
			_, err := fmt.Scanln(&str)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			fmt.Println(second.StringLength(str))
		case 5:
			var width, height int
			fmt.Print("Введите ширину и высоту прямоугольника: ")
			_, err := fmt.Scanln(&width, &height)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			rectangle := second.Rectangle{Width: width, Height: height}
			fmt.Printf("Площадь вашего прямоугольника равна = %d\n", second.Square(rectangle))
		case 6:
			var a, b int
			fmt.Print("Введите два целых числа: ")
			_, err := fmt.Scanln(&a, &b)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			fmt.Printf("Среднее арифмитическое = %f\n", second.Average(a, b))

		default:
			fmt.Println("Неверный номер задания")
		}
	case 3:
		fmt.Print("Введите номер задания (1-6): ")
		_, err := fmt.Scanln(&taskNumber)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		switch taskNumber {
		// 2) Использовать созданный пакет для вычисления факториала введенного пользователем числа.
		case 1, 2:
			var num int
			fmt.Print("Введите число для вычисления факториала ")
			_, err := fmt.Scanln(&num)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
			fmt.Println(mathutils.Factorial(num))
		case 3:
			var str string
			fmt.Print("Введите строку, которую хотите перевернуть: ")
			_, err := fmt.Scanln(&str)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				return
			}
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
	default:
		fmt.Println("Неверный номер лабораторной работы")
	}
}
