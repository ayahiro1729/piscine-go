package main

import (
	"bufio"
	"os"
)

func ft_print_comb() {
	var buffer *bufio.Writer
	var i byte
	var j byte
	var k byte

	buffer = bufio.NewWriter(os.Stdout)
	i = '0'
	for i < '6' {
		j = i + 1
		for j < '8' {
			k = j + 1
			for k < '9' {
				buffer.WriteByte(i)
				buffer.Flush()
				buffer.WriteByte(j)
				buffer.Flush()
				buffer.WriteByte(k)
				buffer.Flush()
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
