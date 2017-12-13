package main

import "fmt"

var number int = 123
var reverse int = 0
var sd int = 0

func main() {
	for ;number > 0; {
		sd = number % 10
		reverse = reverse * 10 + sd
		number = number/10
	}
	fmt.Println(reverse)
}