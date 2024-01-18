package main

import "fmt"

func recursiveFunction(n int) {
	if n > 0 {
		recursiveFunction(n / 2)
		fmt.Println(n)
	}
}

func main() {
	recursiveFunction(9)
}
