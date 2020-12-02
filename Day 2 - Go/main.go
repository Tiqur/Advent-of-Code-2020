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
type Output struct {
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
func helperFunc(input []string) []*Output {
	var formattedInputs []*Output
	for _, inputString := range input {
		var min string
		var max string
		output := new(Output)

		for i := 0; i < len(inputString); i++ {
			char := string(inputString[i])
			hyphenIndex := strings.Index(inputString, "-")
			colonIndex := strings.Index(inputString, ":")

			switch {
			// get min num
			case i < hyphenIndex:
				min += char
			// get max num
			case i < colonIndex - 2 && i > hyphenIndex:
				max += char
			// get character to test
			case i == colonIndex - 1:
				output.character = char
			// get password / string
			case i > colonIndex + 1:
				output.password += char
			}
		}

		minInt, _ := strconv.Atoi(min)
		maxInt, _ := strconv.Atoi(max)
		output.min = minInt
		output.max = maxInt
		formattedInputs = append(formattedInputs, output)
	}
	return formattedInputs
}

func part1(inputs []*Output) int {
	valid := 0
	for _, input := range inputs {
		characters := 0
		for _, character := range input.password {
			if string(character) == input.character {
				characters++
			}
		}
		if characters >= input.min && characters <= input.max {
			valid++
		}
	}
	return valid
}

func part2(inputs []*Output) int {
	valid := 0
	for _, input := range inputs {
		index1char := string(input.password[input.min-1])
		index2char := string(input.password[input.max-1])
		if !(input.character == index1char && input.character == index2char) && (input.character == index1char || input.character == index2char) {
			valid++
		}
	}
	return valid
}