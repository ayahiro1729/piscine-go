package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func rc(prev int, n int, g_length int, g_array *[]byte) {
	var i int

	if n == g_length {
		os.Stdout.Write((*g_array)[:g_length])
		os.Stdout.Write([]byte(", "))
		return
	}
	if prev == 9 {
		return
	}
	i = prev
	for i <= 10-g_length+n {
		(*g_array)[n] = byte('0' + i)
		i++
		rc(i, n + 1, g_length, g_array)
	}
}

func print_last(i int) {
	var c byte

	for i <= 9 {
		c = byte('0' + i)
		ft_putchar(c)
		i++
	}
}

func ft_print_combn(n int) {
	var i int
	var g_length int
	var g_array  []byte

	i = 0
	g_length = n
	g_array = make([]byte, g_length)
	for i < 10-n {
		g_array[0] = byte('0' + i)
		i++
		rc(i, 1, g_length, &g_array)
	}
	print_last(i)
}

func main() {
	ft_print_combn(10)
}
