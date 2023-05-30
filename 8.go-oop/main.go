package main

import (
	"fmt"
	"go-oop/model"
)

type Singer interface {
	Sing() string
}

func main() {

	// inheritance
	g := model.NewGirl("Lady", 25)

	fmt.Println(g.GetName())
	fmt.Println(g.GetAge())

	// polymorphism

	bird := model.Bird{}

	fmt.Println(
		sing(g),
		sing(bird),
	)
}

func sing(singer Singer) string {
	return singer.Sing()
}
