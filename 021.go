package main

import "math"

const N = 10000

func main() {
	var (
		fi, div, sqrt float64
		sum, idiv, isqrt int
		d [N + 1]int
		checked [N + 1]bool
	)
	for i := 1; i <= N; i++ {
		fi = float64(i)
		sqrt = math.Sqrt(fi)
		isqrt = int(sqrt)
		if sqrt == float64(isqrt) { d[i] = isqrt }
		for n := 1; n < isqrt; n++ {
			div = fi / float64(n)
			idiv = int(div)
			if div == float64(idiv) {
				if div < fi { d[i] += idiv }
				d[i] += n
			}
		}
	}
	for i := 1; i <= N; i++ {
		if !checked[i] && d[i] <= N && i != d[i] && i == d[d[i]] {
			checked[d[i]] = true
			sum += i + d[i]
		}
		checked[i] = true
	}
	println(sum)
}