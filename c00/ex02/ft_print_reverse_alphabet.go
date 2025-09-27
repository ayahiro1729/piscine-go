package main

import (
	"bufio"
	"os"
)

func ft_print_alphabet() {
	var buffer *bufio.Writer
	var c byte

	buffer = bufio.NewWriter(os.Stdout)
	c = 'z'
	for c >= 'a' {
		buffer.WriteByte(c)
		buffer.Flush()
		c--
	}
}

func main() {
	ft_print_alphabet()
}
