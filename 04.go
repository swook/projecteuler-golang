package main

import "fmt"
import "strconv"

func main() {
	s, n, p, z := "", 0, 0, 0
	for a := 999; a > 800; a-- {
		L: for b := 999; b > 800; b-- {
			n = a * b
			s = strconv.Itoa(n)
			z = int(len(s)/2)
			for i := 0; i < z; i++ {
				if s[i] != s[len(s) - i - 1] { continue L }
			}
			if p == 0 || n > p { p = n }
		}
	}
	fmt.Println(p)
}

