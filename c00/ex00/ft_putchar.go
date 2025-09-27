package main

import "os"

func ft_putchar(str string) {
	os.Stdout.Write([]byte(str))
}

func  main() {
	ft_putchar("Hello world!")
}
