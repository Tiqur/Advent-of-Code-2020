// GoLang
// Quick disclaimer, I have never used golang before
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// helper function output struct
type Input struct {
	character, password string
	min, max int
}

func main() {
	var inputs []string
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	formattedInputs := helperFunc(inputs)
	fmt.Println("Part 1:", part1(formattedInputs))
	fmt.Println("Part 2:", part2(formattedInputs))
}


// formats inputs so they can be easily solved
func helperFunc(input []string) []*Input {
	var formattedInputs []*Input
	for _, inputString := range input {
		output := new(Input)

		// int 1 and 2
		minInt, _ := strconv.Atoi(strings.Split(inputString, "-")[0])
		maxInt, _ := strconv.Atoi(strings.Split(strings.Split(inputString, "-")[1], " ")[0])
		output.min = minInt
		output.max = maxInt

		// character
		output.character = string(inputString[strings.Index(inputString, ":")-1])

		// password
		output.password = strings.Split(inputString, " ")[2]

		formattedInputs = append(formattedInputs, output)
	}
	return formattedInputs
}

func part1(inputs []*Input) int {
	valid := 0
	for _, input := range inputs {
		characters := strings.Count(input.password, input.character)
		if characters >= input.min && characters <= input.max {
			valid++
		}
	}
	return valid
}

func part2(inputs []*Input) int {
	valid := 0
	for _, input := range inputs {
		index1char := string(input.password[input.min-1])
		index2char := string(input.password[input.max-1])
		notBoth := !(input.character == index1char && input.character == index2char)
		oneOrOther := input.character == index1char || input.character == index2char
		if notBoth && oneOrOther {
			valid++
		}
	}
	return valid
}