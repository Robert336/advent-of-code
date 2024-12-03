package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Objective: analyize reports
	// One report per line, numbers separated by spaces
	// each number in a report is a "level"

	fmt.Println("Advent of Code Day 2: Part 1")

	file, err := os.Open("day2-reports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// to keep track of the total safe reports
	safe_report_count := 0

	// parse file
	for scanner.Scan() {
		current_line := scanner.Text()
		parts := strings.Fields(current_line)

		report := make([]int, 0)
		// store each parsed level as an int within the report slice
		for i := 0; i < len(parts); i++ {
			level, _ := strconv.Atoi(parts[i])
			report = append(report, level)
		}

		// going to evaluate the reports as we read them for memory efficency

		if isSafe(report) {
			safe_report_count += 1
		}
	}

	fmt.Println("Safe Reports: ", safe_report_count)

}

// determines if a report is safe given an array of it's "levels"
// find the "safe" reports:
// safe = gradually increasing or gradually decreasnig:
//   - levels are all increasing or all decreasing (no flat-lining)
//   - levels differ by at least one and at most three
func isSafe(report []int) bool {

	if len(report) < 2 {
		return true
	}

	// a flag that represents if levels are increasing or decreasing
	var is_increasing bool

	// loop through levels -- stop one position early because we are going to look-ahead by one index position
	// set is_increasing on first iteration
	// check is direction changed (increasing or decreasing)
	//	if flat then return unsafe
	// 	if levels changed by greater than 3 return unsafe
	//	if the direction changed return unsafe (example: increasing -> decreasing)
	// check the difference between the levels
	// if loop reaches the end; return safe

	for i := 0; i < len(report)-1; i++ {
		if report[i] == report[i+1] || abs(report[i]-report[i+1]) > 3 {
			return false
		} else if i == 0 {
			is_increasing = report[i] > report[i+1]
		} else if report[i] > report[i+1] && is_increasing {
			return false
		} else if report[i] < report[i+1] && !is_increasing {
			return false
		}
	}

	return true
}

// absoloute value for integers
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
