package main

import "fmt"
import "math"

func main() {
	var n, m, c int
	for i := 1; ; i++ {
		n, m, c = i * (i + 1) / 2, int(math.Sqrt(float64(n))), 0
		for f := 1; f < m; f++ {
			if n % f == 0 { c++ }
		}
		c *= 2
		if m * m == n { c += 1 }
		if c > 500 {
			fmt.Println(n)
			break
		}
	}
}
