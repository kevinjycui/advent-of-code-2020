package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func checkLine(matrix [][]rune, i, j, d int) bool {
	for {
		switch d {
		case 0:
			if i-1 < 0 {
				return false
			}
			i--
		case 1:
			if i+1 >= len(matrix) {
				return false
			}
			i++
		case 2:
			if j-1 < 0 {
				return false
			}
			j--
		case 3:
			if j+1 >= len(matrix[0]) {
				return false
			}
			j++
		case 4:
			if i-1 < 0 || j-1 < 0 {
				return false
			}
			i--
			j--
		case 5:
			if i+1 >= len(matrix) || j-1 < 0 {
				return false
			}
			i++
			j--
		case 6:
			if i+1 >= len(matrix) || j+1 >= len(matrix[0]) {
				return false
			}
			i++
			j++
		case 7:
			if i-1 < 0 || j+1 >= len(matrix[0]) {
				return false
			}
			i--
			j++
		}
		if matrix[i][j] == '#' {
			return true
		} else if matrix[i][j] == 'L' {
			return false
		}
	}
}

func check(matrix [][]rune, i, j int, c rune) int {
	count := 0
	for d:=0; d<8; d++ {
		if checkLine(matrix, i, j, d) {
			count++
		}
	}
	return count
}

func step(matrix [][]rune) {
	newmatrix := make([][]rune, len(matrix))
	for i := range newmatrix {
		newmatrix[i] = make([]rune, len(matrix[0]))
		copy(newmatrix[i], matrix[i])
	}
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			adj := check(matrix, i, j, '#')
			if matrix[i][j] == 'L' && adj == 0 {
				newmatrix[i][j] = '#'
			} else if matrix[i][j] == '#' && adj >= 5 {
				newmatrix[i][j] = 'L'
			}
		}
	}
	copy(matrix, newmatrix)
}

func count(matrix [][]rune, v rune) int {
	count := 0
	for _, r := range matrix {
		for _, c := range r {
			if c == v {
				count++
			}
		}
	}
	return count
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(data), "\n")
	matrix := make([][]rune, len(rows)-1)
	for i, r := range rows {
		if i == len(rows)-1 {
			break
		}
		matrix[i] = make([]rune, len(rows[0]))
		for j, c := range r {
			matrix[i][j] = c
		}
	}
	prevCount := -1
	for {
		curCount := count(matrix, '#')
		if curCount == prevCount {
			fmt.Println(curCount)
			break
		}
		prevCount = curCount
		step(matrix)
	}
}
