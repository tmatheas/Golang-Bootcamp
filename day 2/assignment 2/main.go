package main

import (
	"fmt"
	"os"
)

func main() {
	var x string
	var y string
	fmt.Print("Enter string 1: ")
	fmt.Scan(&x)
	fmt.Print("Enter string 2: ")
	fmt.Scan(&y)
	if len(x) != len(y) {
		fmt.Println("Cannot compare strings of different length !")
		os.Exit(1)
	}
	x_rune := []rune(x)
	y_rune := []rune(y)
	count := 0
	for i := 0; i < len(x); i++ {
		if x_rune[i] != y_rune[i] {
			count++
		}
	}
	fmt.Println("Difference: ", count)
}
