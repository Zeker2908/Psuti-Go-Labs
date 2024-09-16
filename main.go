package main

import (
	"PsutiGoLabs/pkg/labs/first"
	"fmt"
)

func main() {
	var labNumber, taskNumber int
	fmt.Print("Введите номер лабораторной работы: ")
	fmt.Scanln(&labNumber)

	switch labNumber {
	case 1:
		// Выбор задания для первой лабораторной работы
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
			var a, b, c float64
			fmt.Print("Введите три числа с плавающей точкой: ")
			fmt.Scanln(&a, &b, &c)
			first.Avg(a, b, c)
		default:
			fmt.Println("Неверный номер задания")
		}
	default:
		fmt.Println("Неверный номер лабораторной работы")
	}
}
