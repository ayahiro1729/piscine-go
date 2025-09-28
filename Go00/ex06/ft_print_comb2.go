package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func ft_print_comb2() {
	var num1 int
	var num2 int

	num1 = 0
	for num1 <= 97 {
		num2 = num1 + 1
		for num2 <= 99 {
			ft_putchar(byte(num1 / 10 + '0'))
			ft_putchar(byte(num1 % 10 + '0'))
			ft_putchar(' ')
			ft_putchar(byte(num2 / 10 + '0'))
			ft_putchar(byte(num2 % 10 + '0'))
			os.Stdout.Write([]byte(", "))
			num2++
		}
		num1++
	}
	os.Stdout.Write([]byte("98 99"))
}

func main() {
	ft_print_comb2()
}
