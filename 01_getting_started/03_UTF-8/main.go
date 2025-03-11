package main

import "fmt"

func main() {
	for i := 0; i < 127; i++ {
		// %d - decimal 十进制
		// %b - binary 二进制
		// %x - hexadecimal 十六进制
		// %q - quoted string
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
