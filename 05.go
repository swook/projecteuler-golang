package main

import "fmt"

func main() {
	p, t, lcd := 1, [2]int{}, 1
	for c:= 2; c <= 20; c++ {
		t[0], t[1], lcd = p, c, 1
		L: for {
			if t[0] > t[1] {
				t[0] =  t[0] % t[1]
			} else { t[1] = t[1] % t[0] }

			if t[0] == 1 || t[1] == 1 { break L }
			if t[0] * t[1] == 0 {
				if t[0] == 0 {
					lcd = t[1]
				} else { lcd = t[0] }
				break L
			}
		}
		p *= c / lcd
	}
	fmt.Println(p)
}

