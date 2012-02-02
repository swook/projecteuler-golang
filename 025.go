package main

const N = 1000

func main() {
	n1 := make([]int, N, N)
	n2 := make([]int, N, N)
	var a, b []int
	n1[0], n2[0] = 1, 1
	for i := 3; ; i++ {
		if i % 2 == 0 {
			a = n1
			b = n2
		} else {
			a = n2
			b = n1
		}
		for n := 0; n < N; n++ {
			if n > 0 && a[n - 1] > 9 {
				a[n] += int(a[n - 1] / 10)
				a[n - 1] %= 10
			}
			a[n] += b[n]
		}
		if a[N - 1] != 0 {
			println(i)
			break
		}
	}
}