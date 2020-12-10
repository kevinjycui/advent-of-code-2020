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
	dp := make([]int, len(adpt))
	dp[0] = 1
	for i, a := range adpt {
		if i+1 < len(adpt) {
			if adpt[i+1]-a<=3 {
				dp[i+1]+=dp[i]
			}
			if i+2 < len(adpt) {
				if adpt[i+2]-a<=3 {
					dp[i+2]+=dp[i]
				}
				if i+3 < len(adpt) && adpt[i+3]-a<=3 {
					dp[i+3]+=dp[i]
				}
			}
		}
	}
	fmt.Println(dp[len(dp)-1])
}
