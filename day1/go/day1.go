package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	s, v string
}

var pairs = []Pair{
	{"zerone", "01"},
	{"oneight", "18"},
	{"twone", "21"},
	{"threeight", "38"},
	{"fiveight", "58"},
	{"sevenine", "79"},
	{"eightwo", "82"},
	{"eighthree", "83"},
	{"nineight", "98"},
	{"zero", "0"},
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

func part1() {
	file, err := os.Open("../day1_1_input.txt")
	if err != nil {
		log.Fatal("unable to open input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineRunes := []rune(line)

		numRunes := make([]rune, 0)
		for _, c := range lineRunes {
			if unicode.IsNumber(c) {
				numRunes = append(numRunes, c)
				break
			}
		}

		for i := len(lineRunes) - 1; i >= 0; i-- {
			if unicode.IsDigit(lineRunes[i]) {
				numRunes = append(numRunes, lineRunes[i])
				break
			}
		}

		numStr := string(numRunes)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("unable to perform number conversion", err)
		}

		total += int(num)
	}

	fmt.Println(total)
}

func part2() {
	file, err := os.Open("../day1_1_input.txt")
	if err != nil {
		log.Fatal("unable to open input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, p := range pairs {
			line = strings.ReplaceAll(line, p.s, p.v)
		}
		lineRunes := []rune(line)

		numRunes := make([]rune, 0)
		for _, c := range lineRunes {
			if unicode.IsNumber(c) {
				numRunes = append(numRunes, c)
				break
			}
		}

		for i := len(lineRunes) - 1; i >= 0; i-- {
			if unicode.IsDigit(lineRunes[i]) {
				numRunes = append(numRunes, lineRunes[i])
				break
			}
		}

		numStr := string(numRunes)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("unable to perform number conversion", err)
		}

		total += int(num)
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
