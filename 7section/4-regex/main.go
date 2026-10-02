package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	text1 := "Hello world! Welcome to Go"

	regGo, err := regexp.Compile(`Gof`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// === regGo:= regexp.MustCompile(`Go`)
	rProductP := regexp.MustCompile(`P\d+`) //looks like Compile, but panic if th expression not parsable

	fmt.Printf("Text '%s', matches 'Go': %t\n", text1, regGo.MatchString(text1))

	text2 := "Products codes: P123, X342, P789"
	firstProduct := rProductP.FindString(text2)
	fmt.Println(firstProduct)

	allPProducts := rProductP.FindAllString(text2, -1)
	fmt.Printf("%+v\n", allPProducts)

}
