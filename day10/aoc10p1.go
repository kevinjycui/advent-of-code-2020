package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	nums := strings.Split(string(data), "\n")
	adpt := make([]int, len(nums))
	for i, a := range nums {
		adpt[i], _ = strconv.Atoi(a)
	}
	sort.Slice(adpt, func(i, j int) bool { return adpt[i]<adpt[j] })
	c1 := 0
	c3 := 1
	for i:=1; i<len(adpt); i++ {
		if adpt[i]-adpt[i-1] == 1 {
			c1++
		} else if adpt[i]-adpt[i-1] == 3 {
			c3++
		}
	}
	fmt.Println(c1*c3)
}
