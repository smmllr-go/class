package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most fun and interesting code in go!")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in Foo")
}

//
