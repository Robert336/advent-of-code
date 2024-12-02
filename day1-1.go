// Day 1 Challenge 1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1-locations.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	locationIDs_left := make([]int, 0)
	locationIDs_right := make([]int, 0)

	// parse file and create 2 slices for each list of location IDs
	for scanner.Scan() {
		current_line := scanner.Text()
		parts := strings.Fields(current_line)

		fmt.Println("reading:", current_line)
		locationID_left, _ := strconv.Atoi(parts[0])
		locationID_right, _ := strconv.Atoi(parts[1])

		locationIDs_left = append(locationIDs_left, locationID_left)
		locationIDs_right = append(locationIDs_right, locationID_right)
	}
	// sort each list of location IDs and take the difference at each index position
	slices.Sort(locationIDs_left)
	slices.Sort(locationIDs_right)

	total_difference := 0
	for i := 0; i < len(locationIDs_left); i++ {
		diff := locationIDs_left[i] - locationIDs_right[i]

		// absolute value (math.Abs() is for floats)
		if diff < 0 {
			diff = -diff
		}
		fmt.Println("left: ", locationIDs_left[i], "| right: ", locationIDs_right[i], "| diff: ", diff)
		total_difference += diff
	}
	fmt.Println("Total difference: ", total_difference)
}
