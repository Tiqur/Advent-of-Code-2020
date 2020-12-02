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

func part1(input []string ) int {
	valid := 0
	for _, inputString := range input {
		var character string
		var min string
		var max string
		charAmount := 0

		for i := 0; i < len(inputString); i++ {
			char := string(inputString[i])
			hyphenIndex := strings.Index(inputString, "-")
			colonIndex := strings.Index(inputString, ":")

			if i < hyphenIndex { // before the hyphen ( min )
				min += char
			} else if i < colonIndex - 2 && i > hyphenIndex { // before the colon ( max )
				max += char
			} else if i == colonIndex - 1 {
				character = char
			} else if i > colonIndex + 1 { // after the final space ( actual input )
				if char == character {
					charAmount++
				}
			}
		}

		minInt, _ := strconv.Atoi(min)
		maxInt, _ := strconv.Atoi(max)

		if charAmount >= minInt && charAmount <= maxInt {
			valid ++
		}

	}

	return valid
}

func part2(input []string ) int {
	valid := 0
	for _, inputString := range input {
		var character string
		var index1 string
		var index2 string
		var prefixLength int

		for i := 0; i < len(inputString); i++ {
			char := string(inputString[i])
			hyphenIndex := strings.Index(inputString, "-")
			colonIndex := strings.Index(inputString, ":")

			if i < hyphenIndex { // before the hyphen ( min )
				index1 += char
			} else if i < colonIndex - 2 && i > hyphenIndex { // before the colon ( max )
				index2 += char
			} else if i == colonIndex - 1 {
				character = char
				prefixLength = colonIndex + 1
			}

		}

		index1int, _ := strconv.Atoi(index1)
		index2int, _ := strconv.Atoi(index2)
		index1char := string(inputString[index1int+prefixLength])
		index2char := string(inputString[index2int+prefixLength])
		if !(character == index1char && character == index2char) && (character == index1char || character == index2char) {
			valid ++
		}

	}


	return valid
}


func main() {
	var inputs []string
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}


	fmt.Println("Part 1:", part1(inputs))
	fmt.Println("Part 2:", part2(inputs))
}