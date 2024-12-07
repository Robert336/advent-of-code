package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// day 3 fix the shop keeper's computer: ignore the invalid characters to complete the multiplication
// search for all occurances "mul(#,#)" where # is any int

// part 2: keep tracks of do() and don't()
// pretty much just a semaphore or flag that tells us to process a mul() or not.

func main() {
	fmt.Println("Advent of Code: Day 3")

	file, err := os.Open("day3.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	mul_enabled := true

	// parse file
	// looks for 'mul(number,number)' or 'do()' or 'don't()'
	r, _ := regexp.Compile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	for scanner.Scan() {
		current_line := scanner.Text()
		// search current_line for valid "mul" statements
		//compute the mul statements and sum the results
		matches := r.FindAllString(current_line, -1)

		// enabled/disable mul

		// process matching statements
		for _, statement := range matches {
			fmt.Println("Currently on: ", statement)
			switch statement {
			case "do()":
				// enable
				mul_enabled = true
			case "don't()":
				// disable
				mul_enabled = false
			default:
				if mul_enabled {
					fmt.Println("Evaluating: ", statement)
					sum += EvaluateMul(statement)
				}
			}
		}
	}
	fmt.Println("Sum: ", sum)
}

func EvaluateMul(mul_statement string) int {
	r, _ := regexp.Compile(`\d+`)
	matches := r.FindAllString(mul_statement, -1)
	num_1, _ := strconv.Atoi(matches[0])
	num_2, _ := strconv.Atoi(matches[1])

	return num_1 * num_2
}
