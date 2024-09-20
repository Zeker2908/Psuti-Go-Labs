package fifth

import (
	"fmt"
	"math"
)

// Person 1)Создать структуру Person с полями name и age. Реализовать метод для вывода информации о человеке.
type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func (person *Person) GetName() string {
	return person.name
}

func (person *Person) GetAge() int {
	return person.age
}

func (person *Person) SetName(name string) {
	person.name = name
}

func (person *Person) SetAge(age int) {
	person.age = age
}

// Birthday 2) Реализовать метод birthday для структуры Person, который увеличивает возраст на 1 год.
func (person *Person) Birthday() {
	person.age++
}

// String Реализовать метод для вывода информации о человеке
func (person Person) String() string {
	return fmt.Sprintf("%s is %d years old", person.name, person.age)
}

// Circle 3) Создать структуру Circle с полем radius и метод для вычисления площади круга.
type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{radius}
}

func (circle *Circle) GetRadius() int {
	return circle.radius
}

func (circle *Circle) SetRadius(radius int) {
	circle.radius = radius
}

func (circle Circle) String() string {
	return fmt.Sprintf("Circle: radius: %d, ", circle.radius)
}

// Area Вычисление площади круга
func (circle *Circle) Area() float64 {
	if circle.radius <= 0 {
		return 0
	}
	return math.Pow(float64(circle.radius), 2) * math.Pi
}

// Shape 4) Создать интерфейс Shape с методом Area(). Реализовать этот интерфейс для структур Rectangle и Circle.
type Shape interface {
	Area() float64
}

// PrintArea 5)Реализовать функцию, которая принимает срез интерфейсов Shape и выводит площадь каждого объекта.
func PrintArea(shape []Shape) {
	for _, shape := range shape {
		fmt.Println("Area: ", shape.Area())
	}
}

// Stringer 6) Создать интерфейс Stringer и реализовать его для структуры Book, которая хранит информацию о книге.
type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s, Year: %d", b.Title, b.Author, b.Year)
}
