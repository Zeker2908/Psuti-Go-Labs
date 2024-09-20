package fourth

import (
	"fmt"
)

// PeopleMap 1)Написать программу, которая создает карту с именами людей и их возрастами. Добавить нового человека и вывести все записи на экран.
func PeopleMap(name []string, age []int) map[string]int {
	if len(name) != len(age) {
		return nil
	}
	mapped := make(map[string]int)
	for i := 0; i < len(age); i++ {
		mapped[name[i]] = age[i]
	}
	return mapped
}

// PrintMap выводит все записи на экран.
func PrintMap(mapped map[string]int) {
	if mapped == nil {
		fmt.Println("Ошибка: длины списков ключей и значений не совпадают.")
		return
	}
	for name, age := range mapped {
		fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
	}
}

// AddElement добавляет нового человека в карту.
func AddElement(mapped map[string]int, key string, value int) {
	mapped[key] = value
}

// AvgAgePeopleMap 2) Реализовать функцию, которая принимает карту и возвращает средний возраст всех людей в карте.
func AvgAgePeopleMap(mapped map[string]int) float64 {
	var slice []float64
	for _, age := range mapped {
		slice = append(slice, float64(age))
	}
	return AvgSlice(slice)

}
func AvgSlice(slice []float64) float64 {
	var sum float64
	for _, value := range slice {
		sum += value
	}
	return sum / float64(len(slice))
}

// DeleteElement 3) Написать программу, которая удаляет запись из карты по заданному имени.
func DeleteElement(mapped map[string]int, key string) {
	delete(mapped, key)
}

// Sum 5) Написать программу, которая считывает несколько чисел, введенных пользователем, и выводит их сумму.
func Sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// IntReverse возвращает новый срез, содержащий элементы исходного среза в обратном порядке.
func IntReverse(numbers []int) []int {
	var reversed []int
	for i := len(numbers) - 1; i >= 0; i-- {
		reversed = append(reversed, numbers[i])
	}
	return reversed
}
