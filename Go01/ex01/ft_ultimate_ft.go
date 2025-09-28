package main

import "fmt"

func ft_ultimate_ft(nbr *******int) {
	*******nbr = 42
}

func main() {
	var nbr int
	var ptr1 *int
	var ptr2 **int
	var ptr3 ***int
	var ptr4 ****int
	var ptr5 *****int
	var ptr6 ******int
	var ptr7 *******int

	nbr = 24
	ptr1 = &nbr
	ptr2 = &ptr1
	ptr3 = &ptr2
	ptr4 = &ptr3
	ptr5 = &ptr4
	ptr6 = &ptr5
	ptr7 = &ptr6
	fmt.Printf("before ft_ft: %d\n", nbr)
	ft_ultimate_ft(ptr7)
	fmt.Printf("after ft_ft: %d\n", nbr)
}
