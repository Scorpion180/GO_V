package main

import "fmt"

func mayor(args ...int) int {
	mayor := 0
	for _, v := range args {
		if mayor < v {
			mayor = v
		}
	}
	return mayor
}

func main() {
	a := []int{1, 4, 2}
	fmt.Print(mayor(a...))
}
