package main

import "fmt"

func add(x, y int) int {
	return x+y;
}

func main() {
	x := 10;
	y := 11;
	fmt.Println(add(x, y));
}