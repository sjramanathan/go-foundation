package main

import "fmt"

func Div(a, b int) int {
	return a / b
}

func SafeDiv(a, b int) (q int, err error) {

	// This is how we recover from a panic
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("ERROR: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
}

func main() {
	a := 6
	b := 0

	// fmt.Printf("%d / %d = %d", a, b, div(a, b))

	result, err := SafeDiv(a, b)
	fmt.Printf("%d / %d = %d - %v", a, b, result, err)
}
