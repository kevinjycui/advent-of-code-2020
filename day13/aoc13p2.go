package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	l := strings.Split(string(data), "\n")
	busids := strings.Split(l[1], ",")
	n := make([]int, 0)
	a := make([]int, 0)
	nt := 1
	for i, b := range busids {
		if b != "x" {
			id, _ := strconv.Atoi(b)
			a = append(a, id-i)
			n = append(n, id)
			nt *= id
		}
	}
	sol := 0
	for i:=0; i<len(n); i++ {
		u := 1
		nn := nt/n[i]
		for {
			if (nn * u) % n[i] == 1 {
				break
			}
			u++
		}
		sol += a[i] * nn * u
	}
	fmt.Println(sol%nt)
}
