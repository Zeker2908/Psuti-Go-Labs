package third

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateAndPrintArray 4) Написать программу, которая создает массив из 5 целых чисел, заполняет его значениями и выводит их на экран.
func CreateAndPrintArray() {
	var arr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
		fmt.Println(arr[i])
	}
}

// MakeSlice 5) Создать срез из массива и выполнить операции добавления и удаления элементов.
func MakeSlice() {
	arr := [5]int{1, 2, 3, 4, 5}

	slice := arr[:]
	fmt.Println("Оригинальный срез:", slice)

	// Добавление элемента в срез
	slice = append(slice, 6)
	fmt.Println("После добавления 6:", slice)

	// Удаление элемента из среза
	indexToRemove := 2
	if indexToRemove < len(slice) {
		slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	}
	fmt.Println("После удаления элемента с индексом 2:", slice)
}

// FindLongestString 6) Написать программу, которая создает срез из строк и находит самую длинную строку.
func FindLongestString(str ...string) string {
	if len(str) == 0 {
		return ""
	}
	longest := str[0]
	for _, str := range str[1:] {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}
