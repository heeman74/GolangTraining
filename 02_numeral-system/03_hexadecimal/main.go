package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)   // %# give 0x notation of the number.
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)   // %X gives capital letter for Hex and %x gives lowecase letter for Hex
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42) // \t is adding a taps.

}
