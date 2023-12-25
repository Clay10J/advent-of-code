package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2() {
	file, err := os.Open("../day2_input.txt")
	if err != nil {
		log.Fatal("unable to open input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		game_id, err := strconv.Atoi(strings.Split(lineSplit[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		rounds := strings.Split(lineSplit[1], ";")

		if isPossible(rounds) {
			total1 += game_id
		}
		total2 += getPower(rounds)
	}

	fmt.Println(total1)
	fmt.Println(total2)
}

func isPossible(rounds []string) bool {
	MAX_RED := 12
	MAX_GREEN := 13
	MAX_BLUE := 14

	highest_red := 0
	highest_green := 0
	highest_blue := 0

	for _, round := range rounds {
		hand := strings.Split(round, ",")
		for _, pull := range hand {
			pull = strings.Trim(pull, " ")
			cc := strings.Split(pull, " ")
			count, err := strconv.Atoi(cc[0])
			if err != nil {
				log.Fatal(err)
			}
			color := cc[1]
			if color == "red" {
				highest_red = max(highest_red, count)
				if highest_red > MAX_RED {
					return false
				}
			} else if color == "green" {
				highest_green = max(highest_green, count)
				if highest_green > MAX_GREEN {
					return false
				}
			} else {
				highest_blue = max(highest_blue, count)
				if highest_blue > MAX_BLUE {
					return false
				}
			}
		}
	}
	return true
}

func getPower(rounds []string) int {
	highest_red := 0
	highest_green := 0
	highest_blue := 0

	for _, round := range rounds {
		hand := strings.Split(round, ",")
		for _, pull := range hand {
			pull = strings.Trim(pull, " ")
			cc := strings.Split(pull, " ")
			count, err := strconv.Atoi(cc[0])
			if err != nil {
				log.Fatal(err)
			}
			color := cc[1]
			if color == "red" {
				highest_red = max(highest_red, count)
			} else if color == "green" {
				highest_green = max(highest_green, count)
			} else {
				highest_blue = max(highest_blue, count)
			}
		}
	}
	return highest_red * highest_green * highest_blue
}

func main() {
	day2()
}
