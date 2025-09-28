package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func ft_print_numbers() {
	var c byte

	c = '0'
	for c <= '9' {
		ft_putchar(c)
		c++
	}
}

func main() {
	ft_print_numbers()
}
