package main

import "fmt"

func main() {
	// %d - decimal 十进制
	// %b - binary 二进制
	// %#X - hexadecimal 十六进制
	fmt.Printf("%d \t %b \t %x \n", 42, 42, 42)
	// 42       101010          2a

	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
	// 42       101010          0x2a

	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	//42       101010          0X2A
}
