package second

import "fmt"

// Parity 1) Написать программу, которая определяет, является ли введенное пользователем число четным или нечетным.
func Parity(num int) bool {
	return num%2 == 0
}

// CheckNumberSign 2)Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".
func CheckNumberSign(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

// PrintNumbers 3) Написать программу, которая выводит все числа от 1 до 10 с помощью цикла for.
func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// StringLength 4) Написать функцию, которая принимает строку и возвращает ее длину.
func StringLength(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

// Rectangle 5) Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.
type Rectangle struct {
	Width  int
	Height int
}

func Square(rectangle Rectangle) int {
	return rectangle.Width * rectangle.Height
}

// Average 6) Написать функцию, которая принимает два целых числа и возвращает их среднее значение.
func Average(a, b int) float64 {
	return float64(a+b) / 2
}
