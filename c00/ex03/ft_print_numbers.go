package main

import (
	"bufio"
	"os"
)

func ft_print_numbers() {
	var buffer *bufio.Writer
	var c byte

	buffer = bufio.NewWriter(os.Stdout)
	c = '0'
	for c <= '9' {
		buffer.WriteByte(c)
		buffer.Flush()
		c++
	}
}

func main() {
	ft_print_numbers()
}
