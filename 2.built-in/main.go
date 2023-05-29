package main

import (
	"fmt"
	"os"
)

type A struct {
	a string
	b string
}

func main() {
	var myStruct = A{
		a: "a",
		b: "b",
	}

	fmt.Println(myStruct)

	var arr [2]int
	fmt.Println(arr[0], arr[1])
	arr[0] = 5
	fmt.Println(arr[0], arr[1])

	var slice = make([]string, 0)
	slice = append(slice, "arr0")

	fmt.Println(slice[0])

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recover panic: ", r)
		}
	}()

	fmt.Println(slice[4])

	fmt.Println("below this line will not run")
}

func openFile() {
	f, err := os.Open("filename.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// do things
}
