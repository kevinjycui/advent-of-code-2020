package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func contains(arr *[25]int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

func hassum(arr *[25]int, val int) bool {
	for _, a := range arr {
		if a*2 != val && val-a >= 0 && contains(arr, val-a) {
			return true
		}
	}
	return false
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	nums := strings.Split(string(data), "\n")
	preamble := new([25]int)
	for i := 0; i<25; i++ {
		preamble[i], _ = strconv.Atoi(nums[i])
	}
	cur := 25
	for {
		if cur >= len(nums)-1 {
			break
		}
		cnum, _ := strconv.Atoi(nums[cur])
		if !hassum(preamble, cnum) {
			fmt.Println(cnum)
		}
		preamble[cur%25] = cnum
		cur++
	}
}
