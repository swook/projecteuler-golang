package main

import "fmt"

func get_chain_len (s int64) int64 {
	var iter int64 = 1
	for s > 1 {
		if s % 2 == 0 {
			s /= 2
		} else {
			s = 3 * s + 1
		}
		iter++
	}
	return iter
}

func main() {
	var i, iter, starting_num, chain_len int64
	for i = 1e6 - 1; i > 5e5; i -= 2 {
		iter = get_chain_len(i)
		if iter > chain_len {
			starting_num = i
			chain_len = iter
		}
	}
	fmt.Println(starting_num)
}
