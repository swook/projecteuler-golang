package main

import (
	"io/ioutil"
	"strings"
	"sort"
)

func main() {
	content, err := ioutil.ReadFile("022.txt")
	if err != nil {
		println("Error reading file")
		return
	}
	names := strings.Split(strings.Replace(string(content), "\"", "", -1), ",")
	content = nil
	sort.Strings(names)

	var sum, score int
	for i := 0; i < len(names); i++ {
		score = 0
		for n := 0; n < len(names[i]); n++ {
			score += int(names[i][n]) - 64
		}
		sum += score * (i + 1)
	}
	println(sum)
}
