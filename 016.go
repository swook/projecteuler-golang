package main

import "fmt"

const Pow = 1000

func main() {
	var n = [Pow]int{}
	n[0] = 1
	for i := 0; i < Pow; i++ {
		for ii := 0; ii < len(n); ii++ {
			if n[ii] > 0 {
				n[ii] = n[ii] << 1
			}
			if ii > 0 && n[ii - 1] >= 10 {
				n[ii] += int(n[ii - 1] / 10)
				n[ii - 1] = n[ii - 1] % 10
			}
		}
	}
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	fmt.Println(sum)
}
