package main

import "fmt"

func main() {
	n := 999999
	max := 9
	len := max + 1
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		nums[i] = i
	}
	
	var fact, dig, dig_ int
	number := make([]int, len)
	for i := max; i >= 0; i-- {
		fact = factorial(i)
		fmt.Println("fact",fact)
		dig_ = int(n/fact)
		dig = nums[dig_]
		number[max-i] = dig
		fmt.Println("dig",dig)
		nums = append(nums[:dig_], nums[dig_+1:]...)
		n -= fact*dig_
		fmt.Println("n",n)
	}
	fmt.Print("\nAnswer: ")
	for _, d := range number {
		fmt.Print(d)
	}
}

func factorial(i int)int {
	if i < 2 {
		return 1
	}
	return i * factorial(i-1)
}
