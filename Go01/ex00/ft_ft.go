package main

import "fmt"

func ft_ft(nbr *int) {
	*nbr = 42
}

func main() {
	var nbr int

	nbr = 24
	fmt.Printf("before ft_ft: %d\n", nbr)
	ft_ft(&nbr)
	fmt.Printf("after ft_ft: %d\n", nbr)
}
