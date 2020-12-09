package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	cmds := strings.Split(string(data), "\n")
	i := 0
	acc := 0
	visited := make([]bool, len(cmds))
	for {
		if visited[i] {
			fmt.Println(acc)
			break
		}
		parts := strings.Split(cmds[i], " ")
		number, _ := strconv.Atoi(parts[1])
		visited[i] = true
		switch parts[0] {
		case "acc":
			acc += number
			i++
		case "jmp":
			i += number
		case "nop":
			i++
		}
	}
}
