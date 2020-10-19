package main

import "fmt"

func intercambiar(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	intercambiar(&a, &b)
	fmt.Println(a, b)
}
