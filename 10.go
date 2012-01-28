package main

import "fmt"
import "math"

func isPrime(x int64)bool {
	if x == 2 { return true }
	if x % 2 == 0 { return false }
	
	var i, z int64 = 3, int64(math.Sqrt(float64(x)))
	for ; i <= z; i += 2 {
		if x % i == 0 { return false }
	}
	return true
}

func main() {
	var i, sum int64 = 3, 2
	for ; i < 2e6; i += 2 {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

