package main

import "fmt"

type MyString string

type Number interface {
	int | float64 | float32 | MyString
	fmt.Stringer
}

func (ms MyString) String() string {
	return string(ms)
}

func Sum[T Number](numbers ...T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// v := Sum("Jane", "Mark") --bad ides

	v := Sum[MyString]("Jane", "Mark")
	fmt.Printf("%T\n", v)
}
