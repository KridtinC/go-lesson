package main

import "fmt"

type myError struct {
	desc string
}

func (m myError) Error() string {
	return m.desc
}

func main() {
	err := doSomething(true)
	fmt.Println(err)
}

func doSomething(returnErr bool) error {
	if returnErr {
		var myErr = myError{
			desc: "this is error",
		}
		return myErr
	}
	return nil
}
