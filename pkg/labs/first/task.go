package first

import (
	"fmt"
	"time"
)

// WhatTimeIsIt 1) Написать программу, которая выводит текущее время и дату.
func WhatTimeIsIt() {
	fmt.Println("Текущее время: ", time.Now().Format("2006-01-02 15:04:05"))
}

// PrintOfNumbers 2) Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
// 3) Использовать краткую форму объявления переменных для создания и вывода переменных.
func PrintOfNumbers() {
	var numInt1 int = 10
	numInt2 := 11
	fmt.Println("Целые числа: ", numInt1, numInt2)

	var numFloat1 float64 = 1.1
	numFloat2 := 1.2
	fmt.Println("Вещественные числа: ", numFloat1, numFloat2)

	var string1 string = "Строковый литерал"
	string2 := "Литерал строковый"
	fmt.Println("Строки: ", string1, string2)

	var bool1 bool = true
	bool2 := false
	fmt.Println("Логический тип данный: ", bool1, bool2)
}

// CalculateInt 4) Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
func CalculateInt(a, b int) {
	fmt.Printf("Результат операции с числами %d и %d\n", a, b)
	fmt.Printf("Сложение: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Вычитание: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, a*b)
	if b == 0 {
		fmt.Printf("На ноль делить нельзя!")
	} else {
		fmt.Printf("Деление: %d / %d = %d\n", a, b, a/b)
		fmt.Printf("Остаток от деления: %d %% %d = %d\n", a, b, a%b)
	}

}

// CalculateFloat 5) Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
func CalculateFloat(a, b float64) {
	fmt.Printf("Результат операции с числами %f и %f\n", a, b)
	fmt.Printf("Сложение: %f + %f = %f\n", a, b, a+b)
	fmt.Printf("Вычитание: %f - %f = %f\n", a, b, a-b)
}

// Avg 6) Написать программу, которая вычисляет среднее значение трех чисел.
func Avg(a float64, b float64, c float64) {
	fmt.Printf("Среднее арифмитическое %f\n", (a+b+c)/3)
}
