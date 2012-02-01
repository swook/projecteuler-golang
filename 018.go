package main

import (
	"io/ioutil"
	"flag"
	"strings"
	"strconv"
)

func main() {
	flag.Parse()
	ifile := flag.Arg(0)
	if len(ifile) == 0 {
		println("No File Specified")
		return
	}
	content, err := ioutil.ReadFile(ifile)
	if err != nil {
		println("Error reading", ifile)
		return
	}
	content_arr := strings.Split(string(content), "\n")
	content = nil
	lines := 0
	for i := len(content_arr) - 1; i > 0; i-- {
		if lines == 0 && content_arr[i] != "" {
			lines = i + 1
			break
		}
	}

	var strs []string
	data := make([][]int, lines)
	for i := 0; i < lines; i++ {
		strs = strings.Fields(content_arr[i])
		data[i] = make([]int, len(strs))
		for ii := 0; ii < len(strs); ii++ {
			data[i][ii], _ = strconv.Atoi(strs[ii])
		}
	}
	content_arr = nil
	strs = nil

	trisums := make([][]int, lines)
	for i := lines - 1; i > -1; i-- {
		trisums[i] = make([]int, i + 1)
		for ii := 0; ii < i + 1; ii++ {
			if i == lines - 1 {
				trisums[i][ii] = data[i][ii]
			} else {
				if trisums[i + 1][ii] > trisums[i + 1][ii + 1] {
					trisums[i][ii] = data[i][ii] + trisums[i + 1][ii]
				} else {
					trisums[i][ii] = data[i][ii] + trisums[i + 1][ii + 1]
				}
			}
		}
	}
	println(trisums[0][0])
}
