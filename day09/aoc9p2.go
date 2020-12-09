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
	inv := 0
	for {
		if cur >= len(nums)-1 {
			break
		}
		cnum, _ := strconv.Atoi(nums[cur])
		if !hassum(preamble, cnum) {
			inv = cnum
			break
		}
		preamble[cur%25] = cnum
		cur++
	}
	left := 0
	right := 1
	lnum, _ := strconv.Atoi(nums[0])
	rnum, _ := strconv.Atoi(nums[1])
	sum := lnum + rnum
	for {
		if right >= len(nums)-1 {
                        break
                }
		lnum, _ = strconv.Atoi(nums[left])
		rnum, _ = strconv.Atoi(nums[right+1])
		if sum < inv {
			sum += rnum
			right++
		} else if sum > inv && left < right-1 {
			sum -= lnum
			left++
		} else if sum == inv {
			min := int(^uint(0) >> 1)
			max := -1
			for i := left; i<=right; i++ {
				tnum, _ := strconv.Atoi(nums[i])
				if tnum > max {
					max = tnum
				}
				if tnum < min {
					min = tnum
				}
			}
			fmt.Println(min + max)
			break
		} else {
			sum += rnum
			right++
		}
	}
}
