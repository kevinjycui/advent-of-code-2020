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
	dep, _ := strconv.Atoi(l[0])
	busids := strings.Split(l[1], ",")
	min := int(^uint(0) >> 1)
	sol := -1
	for _, id := range busids {
		if id != "x" {
			nid, _ := strconv.Atoi(id)
			if (nid - dep%nid < min) {
				min = nid - dep%nid
				sol = min * nid
			}
		}
	}
	fmt.Println(sol)
}
