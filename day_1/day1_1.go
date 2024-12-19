package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Main idea here is to sort both slices and find absolute 
// difference between elements. When we sort in ascending order - we get
// smallest values in front of slices

func main() {

	// Read input file:
	left, right := ReadFileToSlices("myInput.txt")
	
	// This is the core of the idea:
	totalDistance := findTotalDistance(left, right)

	// Just print the result:
	fmt.Println("Total distance is: ", totalDistance)

	similarity := FindSimilarity(left, right)

	fmt.Println("Total similarity is: ", similarity)
}

func findTotalDistance(left, right []int) int {
	distance := 0
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := range left {
		distance += abs(left[i] - right[i])
	}


	return distance
}

func ReadFileToSlices(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Problem, reading file:", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Put values in two slices:
	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), "   ")

		leftNum, err := strconv.Atoi(str[0])
		if err != nil {
			fmt.Println("Problem with atoi to left:", err)
		}
		rightNum, err := strconv.Atoi(str[1])
		if err != nil {
			fmt.Println("Problem with atoi to right:", err)
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
		
	}
	return left, right
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}