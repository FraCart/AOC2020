package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var countFirstPolicy int
	var countSecondPolicy int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		a := regexp.MustCompile(`\-|\s|\:`).Split(scanner.Text(), -1)
		posMin, _ := strconv.Atoi(a[0])
		posMax, _ := strconv.Atoi(a[1])
		value := []byte(a[2])
		password := a[4]

		if (password[posMin-1] == value[0]) != (password[posMax-1] == value[0]) {
			countSecondPolicy++
		}
	}

	for scanner.Scan() {
		a := regexp.MustCompile(`\-|\s|\:`).Split(scanner.Text(), -1)
		freqMin, _ := strconv.Atoi(a[0])
		freqMax, _ := strconv.Atoi(a[1])
		value := a[2]
		password := a[4]

		res := strings.Count(password, value)

		if res >= freqMin && res <= freqMax {
			countFirstPolicy++
		}
	}

	fmt.Println(countSecondPolicy)
}
