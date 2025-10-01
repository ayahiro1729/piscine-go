package main

import "fmt"

func ft_swap(a *int, b *int) {
	var tmp int

	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	var a int
	var b int

	a = 42
	b = 24
	fmt.Printf("before swap: a=%d, b=%d\n", a, b)
	ft_swap(&a, &b)
	fmt.Printf("after swap: a=%d, b=%d\n", a, b)
}