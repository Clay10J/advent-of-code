package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
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
