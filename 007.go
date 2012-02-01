package main

import "fmt"
import "math"

func isPrime(x int64)bool {
	if x == 2 { return true }
	var i, z int64 = 2, int64(math.Sqrt(float64(x)))
	for ; i <= z; i++ {
		if x % i == 0 { return false }
	}
	return true
}

func main() {
	var i int64 = 3
	n := 1
	for {
		if isPrime(i) { n++ }
		if n == 10001 {
			fmt.Println(i)
			break
		}
		i++
	}
}

