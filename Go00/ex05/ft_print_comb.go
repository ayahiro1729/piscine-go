package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func ft_print_comb() {
	var i byte
	var j byte
	var k byte

	i = '0'
	for i < '6' {
		j = i + 1
		for j < '8' {
			k = j + 1
			for k < '9' {
				ft_putchar(i)
				ft_putchar(j)
				ft_putchar(k)
				os.Stdout.Write([]byte(", "))
				k++
			}
			j++
		}
		i++
	}
	os.Stdout.Write([]byte("789"))
}

func main() {
	ft_print_comb()
}
