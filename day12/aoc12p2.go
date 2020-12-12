package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	instr := strings.Split(string(data), "\n")
	x := 0
	y := 0
	wx := -10
	wy := 1
	for j, i := range instr {
		if j == len(instr)-1 {
			break
		}
		ch, _ := strconv.Atoi(i[1:len(i)])
		switch rune(i[0]) {
		case 'L':
			for k:=0; k<ch/90; k++ {
				twy := wy
				wy = -wx
				wx = twy
			}
		case 'R':
			for k:=0; k<ch/90; k++ {
				twx := wx
				wx = -wy
				wy = twx
			}
		case 'F':
			x += wx*ch
			y += wy*ch
		case 'E':
			wx -= ch
		case 'W':
			wx += ch
		case 'N':
			wy += ch
		case 'S':
			wy -= ch
		}
	}
	fmt.Println(abs(x)+abs(y))
}
