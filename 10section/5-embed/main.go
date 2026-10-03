package main

import (
	"embed"
	"fmt"
	"log"
)

// enterprise application in Go
// ----------------------
var name = "Joseph"

//go:embed public
var public embed.FS

//go:embed public
var public2 embed.FS

func main() {

	data, err := public.ReadFile("public/data.txt")
	fmt.Println(data)
	data, err = public2.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// public.ReadFile("data.txt") bad idea!
	fmt.Println(data)

	fmt.Println(string(data))
}
