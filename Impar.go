package main

import "fmt"

func generadorImpar() func() uint {
	i := uint(1)
	return func() uint {
		var par = i
		i += 2
		return par
	}
}

func main() {
	nextImpar := generadorImpar()
	fmt.Println(nextImpar())

	fmt.Println(nextImpar())

	fmt.Println(nextImpar())

	fmt.Println(nextImpar())
}
