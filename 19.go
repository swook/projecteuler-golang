package main

func main() {
	count, month, day := 0, 1, 1
	for year := 1900; year < 2001; year++ {
		for month = 1; month < 13; month++ {
			switch month {
				case 1, 3, 5, 7, 8, 10, 12: day += 31
				case 4, 6, 9, 11: day += 30
				case 2: {
					if year % 4 == 0 {
						if year % 100 == 0 {
							if year % 400 == 0 { day += 29 }
						} else { day += 29 }
					}
				}
			}
			day = day % 7
			if day == 0 && year > 1900 { count++ }
		}
	}
	println(count)
}