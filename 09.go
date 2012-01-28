package main

import "fmt"

const N = 50

func main() {
	k, m, n := 1, N, N
	a, b, c, p := 0, 0, 0, 0
	L: for ; m > 1; m-- {
		n = m - 1
		for ; n > 1; n-- {
			k = 1
			for ; k < 30; k++ {
				if 2 * k * m * (m + n) == 1000 {
					a, b, c = k * (m * m - n * n), 2 * k * m * n, k * (m * m + n * n)
					if a * b * c > 0 {
						p = a * b * c
						break L
					}
				}
			}
		}
	}
	fmt.Println(p)
}

