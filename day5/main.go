package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// TotalSeats ...
// const TotalSeats = 127

func main() {
	input := readInput()
	fmt.Println("Max value:", findMax(input))
	fmt.Println("My seat:", findSeat(input))
}

func readInput() []int {
	data, readFileErr := ioutil.ReadFile("input.txt")
	if readFileErr != nil {
		panic(readFileErr)
	}
	lines := strings.Split(string(data), "\n")
	var res []int
	replacer := strings.NewReplacer("B", "1", "R", "1", "F", "0", "L", "0")
	for _, line := range lines {
		tmp, _ := strconv.ParseInt(replacer.Replace(line), 2, 0)
		i0 := int(tmp)
		res = append(res, i0)
	}
	return res
}

func findMax(input []int) (max int) {
	max = input[0]
	for _, value := range input {
		if value > max {
			max = value
		}
	}
	return max
}

func findSeat(input []int) (seat int) {
	seat = input[0]
	sort.Ints(input)
	for i, value := range input {
		if i < len(input)-1 {
			if input[i+1]-value == 2 {
				seat = value + 1
			}
		}
	}
	return seat
}
