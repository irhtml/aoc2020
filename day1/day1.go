package main

import (
	"fmt"
	"github.com/irhtml/aoc2020/lib"
	"github.com/irhtml/aoc2020/lib/inputs"
)

var Input = lib.InputInts(inputs.Day1(), lib.NumberParser)


func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

func part1(input [][]int) (rc int) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i][0]+input[j][0] == 2020 {
				return input[i][0] * input[j][0]
			}
		}
	}
	return
}

func part2(input [][]int) (rc int) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i][0] + input[j][0] + input[k][0] == 2020 {
					return input[i][0] * input[j][0] * input[k][0]
				}
			}
		}
	}
	return
}





