package main

import (
	"fmt"
	"strconv"
)

var (
	domain = []string{"thegeeklinux", "rajugupta"}
)

func main() {
	for i := 0; i < 12; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(domain[j] + strconv.Itoa(i+1))
		}
	}
}
