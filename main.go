package main

import "fmt"

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	i := increment()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
}
