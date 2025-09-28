package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func ft_print_alphabet() {
	var c byte

	c = 'z'
	for c >= 'a' {
		ft_putchar(c)
		c--
	}
}

func main() {
	ft_print_alphabet()
}
