package main

import "fmt"
import "math"

const n int64 = 600851475143

func isPrime(x int64)bool {
	var i int64 = 2
	for ; i < x; i++ {
		if x % i == 0 { return false }
	}
	return true
}

func main() {
	var i int64 = int64(math.Sqrt(float64(n)))
	for ; i > 1; i-- {
		if n % i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}

