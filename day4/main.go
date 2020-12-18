package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

	firstAnswer := firstValidation(input, validFields)
	secondAnswer := secondValidation(input, validFields)

	fmt.Println("First answer:", firstAnswer)
	fmt.Println("Second answer:", secondAnswer)
}

func readInput() []string {
	data, readFileErr := ioutil.ReadFile("input.txt")
	if readFileErr != nil {
		panic(readFileErr)
	}
	lines := strings.Split(string(data), "\n\n")
	return lines
}

func firstValidation(input []string, validFields []string) int {
	count := 0
	for _, passport := range input {
		tmp := 0
		fields := strings.Fields(passport)
		for _, validField := range validFields {
			for _, field := range fields {
				if strings.Contains(field, validField) || validField == "cid" {
					tmp++
					break
				}
			}
		}
		if tmp == len(validFields) {
			count++
		}
	}
	return count
}

func secondValidation(input []string, validFields []string) int {
	count := 0
	for _, passport := range input {
		tmp := 0
		fields := strings.Fields(passport)
		for _, validField := range validFields {
			for _, field := range fields {
				if strings.Contains(field, validField) || validField == "cid" {
					if fieldValidation(field) {
						tmp++
					}
					break
				}
			}
		}
		if tmp == len(validFields) {
			count++
		}
	}
	return count
}

func fieldValidation(input string) bool {
	r := false
	switch input[0:3] {
	case "byr":
		date, _ := strconv.Atoi(input[4:])
		if date >= 1920 && date <= 2002 {
			r = true
		}
	case "iyr":
		date, _ := strconv.Atoi(input[4:])
		if date >= 2010 && date <= 2020 {
			r = true
		}
	case "eyr":
		date, _ := strconv.Atoi(input[4:])
		if date >= 2020 && date <= 2030 {
			r = true
		}
	case "hgt":
		value := []string{input[4 : len(input)-2], input[len(input)-2:]}
		height, _ := strconv.Atoi(value[0])
		unit := value[1]
		switch unit {
		case "cm":
			if height >= 150 && height <= 193 {
				r = true
			}
		case "in":
			if height >= 59 && height <= 76 {
				r = true
			}
		}
	case "hcl":
		if strings.Contains(input[4:], "#") {
			a := regexp.MustCompile("^[a-f0-9]*$")
			if a.MatchString(input[5:]) {
				if len(input[5:]) == 6 {
					r = true
				}
			}
		}
	case "ecl":
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, color := range validColors {
			if strings.Contains(input[4:], color) {
				r = true
				break
			}
		}
	case "pid":
		a := regexp.MustCompile("^[0-9]*$")
		if a.MatchString(input[4:]) {
			if len(input[4:]) == 9 {
				r = true
			}
		}
	case "cid":
		r = true
	}
	return r
}
