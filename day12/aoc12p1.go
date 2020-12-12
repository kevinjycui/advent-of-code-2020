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
	dirs := [4]rune{'E', 'N', 'W', 'S'}
	fac := 0
	x := 0
	y := 0
	for j, i := range instr {
		if j == len(instr)-1 {
			break
		}
		ch, _ := strconv.Atoi(i[1:len(i)])
		dir := rune(i[0])
		if dir == 'L' {
			fac = (fac+(ch/90))%4
		} else if dir == 'R' {
			fac = (fac-(ch/90)+4)%4
		} else if dir == 'F' {
			dir = dirs[fac]
		}
		if dir == 'E' {
			x -= ch
		} else if dir == 'W' {
			x += ch
		} else if dir == 'N' {
			y += ch
		} else if dir == 'S' {
			y -= ch
		}
	}
	fmt.Println(abs(x)+abs(y))
}
