package main

import "fmt"

const N = 100

func main() {
	ans := N * (N + 1) / 2
	ans *= ans
	sos := 0
	for i := 1; i <= N; i++ { sos += i * i }
	ans -= sos
	fmt.Println(ans)
}

