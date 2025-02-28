package main

import "fmt"

func main() {
	s := "Hello, World!"
	x := s[1]
	y := string(x)
	// それぞれのデータ型と値を出力
	fmt.Printf("x: %T, %v\n", x, x) // x: uint8, 101
	fmt.Printf("y: %T, %v\n", y, y) // y: string, e
}
