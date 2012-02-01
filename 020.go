package main

const N = 100

func main() {
	num := [200]int{}
	num[0] = 1
	for i := 2; i <= N; i++ {
		for a := 0; a < len(num); a++ {
			num[a] *= i
			if a > 0 && num[a - 1] > 9 {
				num[a] += int(num[a - 1] / 10)
				num[a - 1] %= 10
			}
		}
	}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	println(sum)
}