package main

import (
	"bufio"
	"os"
)

func ft_print_alphabet() {
	var buffer *bufio.Writer
	var c byte

	buffer = bufio.NewWriter(os.Stdout)
	c = 'a'
	for c <= 'z' {
		buffer.WriteByte(c)
		buffer.Flush()
		c++
	}
}

func main() {
	ft_print_alphabet()
}
