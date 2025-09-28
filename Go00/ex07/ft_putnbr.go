package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func ft_putnbr(nb int) {
	if nb == -2147483648 {
		os.Stdout.Write([]byte("-2147483648"))
		return
	}
	if nb < 0 {
		nb *= -1
		ft_putchar('-')
	}
	if nb >= 10 {
		ft_putnbr(nb / 10)
	}
	ft_putchar(byte(nb % 10 + '0'))
}

func main() {
	ft_putnbr(42)
	ft_putchar('\n')
	ft_putnbr(-42)
	ft_putchar('\n')
	ft_putnbr(-2147483648)
	ft_putchar('\n')
	ft_putnbr(2147483647)
	ft_putchar('\n')
	ft_putnbr(0)
}
