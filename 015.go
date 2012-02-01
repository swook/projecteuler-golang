package main

import "fmt"

func main() {
	var i, prod float64 = 0, 1
	for i = 40; i > 20; i-- {
		prod *= i / (i - 20)
		fmt.Println(i/(i-20))
	}
	fmt.Println(prod)
}
