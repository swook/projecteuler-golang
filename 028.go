package main

const LIMIT = 1001

func main() {
	i, s, w, sum := 1, 0, 1, 1
	for {
		for n := 0; n < 5; n++ {
			if n < 2 { s++ }
			if n == 0 {
				w += 2
			} else {
				i += s
				sum += i
			}
		}
		if w == LIMIT { break }
	}
	println(sum)
}