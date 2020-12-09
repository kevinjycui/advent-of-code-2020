package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(j int, acc int, cmds []string, visited []bool) (bool, int) {
	iparts := strings.Split(cmds[j], " ")
	icmd := iparts[0]
	inum, _ := strconv.Atoi(iparts[1])
	if icmd == "jmp" {
		j++
	} else if icmd == "nop" {
		j += inum
	}
	nvisited := make([]bool, len(cmds))
	_ = copy(nvisited, visited)
	for {
		if nvisited[j] {
			return false, 0
		}
		parts := strings.Split(cmds[j], " ")
		cmd := parts[0]
		number, _ := strconv.Atoi(parts[1])
		nvisited[j] = true
		switch cmd {
		case "acc":
			acc += number
			j++
		case "jmp":
			j += number
		case "nop":
			j++
		}
		if j == len(cmds)-1 {
			return true, acc
		} else if j >= len(cmds) {
			return false, 0
		}
	}
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	cmds := strings.Split(string(data), "\n")
	i := 0
	acc := 0
	visited := make([]bool, len(cmds))
	for {
		if visited[i] {
			break
		}
		parts := strings.Split(cmds[i], " ")
		number, _ := strconv.Atoi(parts[1])
		visited[i] = true
		if parts[0] != "acc" {
			fix, nacc := check(i, acc, cmds, visited)
			if fix {
				fmt.Println(nacc)
			}
		}
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
