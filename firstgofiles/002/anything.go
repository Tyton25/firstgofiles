package main

import "fmt"

func main() {
	fmt.Println("Running functions 'foo' & 'bar'.")

	foo()

	fmt.Println("Something more.")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("In the foo function.")
}

func bar() {
	fmt.Println("bar exits anything.go")
}
