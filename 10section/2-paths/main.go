package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	path1 := filepath.Join("C:", "Users", "Documents")
	fmt.Println(path1)

	path2 := filepath.Join("/home", "cas12")
	fmt.Println(filepath.Dir(path2))

	fmt.Println(filepath.Base(path2))
	fmt.Println("Is Exists", filepath.Ext("/home/cas12/vscode/learning-go/9section/7-downloader/main.go"))

	dirtyDir := "../users/./dir/../other_dir/./file.txt"
	fmt.Println(filepath.Clean(dirtyDir))

}
