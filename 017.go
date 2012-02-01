package main

import "fmt"

var (
	// nil, one, two, three, four, five, six, seven, eight, nine
	sub10lc = [10]int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4}

	// nil, ten, twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety
	sub100lc = [10]int{0, 3, 6, 6, 5, 5, 5, 7, 6, 6}

	// ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen
	teenlc = [10]int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
)

func main() {
	d, sum := 0, 0
	for i := 0; i <= 1000; i++ {
		d = i % 100
		switch {
		case i == 1000: sum += 11
		case i > 99: {
			sum += sub10lc[int(i / 100)] + 7
			if d > 0 { sum += 3 }
			fallthrough
		}
		default: {
			switch {
				case d > 19: sum += sub100lc[int(d / 10)]
				case d > 9: sum += teenlc[d - 10]
			}
			if d < 10 || d > 20 { sum += sub10lc[i % 10] }
		}
		}
		fmt.Printf("Added %d: %d\n", i, sum)
	}
	fmt.Println(sum)
}
