package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	defer fmt.Println("Function simpleDefer: deferred")
}

func lifoSimpleDefer() {
	fmt.Println("Function lifoSimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Function lifoSimpleDefer: Middle")
}
func main() {
	file, err := os.Create("./defer.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")

}
