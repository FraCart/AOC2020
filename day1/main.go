package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getFile()

	var res int
	for _, v := range input {
		for _, v1 := range input {
			for _, v2 := range input {
				r := v + v1 + v2
				if r == 2020 {
					res = v * v1 * v2
				}
			}
		}
	}

	fmt.Println(res)
}

func getFile() []int {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var slice []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s := scanner.Text()
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, n)
	}

	return slice
}
