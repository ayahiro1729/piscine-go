package main

import "os"

func ft_is_negative(a int) {
	if (a < 0) {
		os.Stdout.Write([]byte("N"))
	} else {
		os.Stdout.Write([]byte("P"))
	}
}

func main() {
	ft_is_negative(42)
	ft_is_negative(-42)
	ft_is_negative(0)
}
