package main

import "fmt"

func main() {
	var p *int
	i := 42
	p = &i
	fmt.Println(p, *p)
	*p = 21
	fmt.Println(p, *p)
	var q = p
	*q = 30
	fmt.Printf("q addr: %p\nq val: %d\np addr: %p\np val: %d\nval: %d\n", q, *q, p, *p, i)
}
