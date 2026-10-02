package main

import "fmt"

func main() {

	names := []string{"Alice", "John", "Mark"}
	fmt.Println(names)
	names = append(names, "Alice")

	items := make([]int, 3, 5)
	fmt.Println(names, cap(items))

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	fmt.Printf("%+v", items[3:7])

}
