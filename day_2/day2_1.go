package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reports := ReadFileAndFindLevels("input.txt")
	safeCount := 0
	safeCountWithout := 0

	for _, report := range reports {
		if isSafeReports(report) {
			safeCount += 1
		} else if isSafeReportWithout(report) {
			safeCountWithout += 1
		}
	}

	fmt.Println("Number of safe reports is:", safeCount)

	fmt.Println("Number of safe reports without one problem is:", safeCountWithout + safeCount)
}

func ReadFileAndFindLevels(filename string) [][]int {
	reports := make([][]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		levels := make([]int, 0, len(str))
		for _, val := range str {
			level, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error converting ASCII to int:", err)
				return nil
			}
			levels = append(levels, level)
		}
		fmt.Println(levels)

		reports = append(reports, levels)
	}
	return reports
}

func isSafeReports(report []int) bool {
	if sort.SliceIsSorted(report, func(i, j int) bool {
		return report[i] < report[j]
	}) || sort.SliceIsSorted(report, func(i, j int) bool {
		return report[i] > report[j]
	}) {
		fmt.Println("This level is sorted", report)
		for i := range report {
			if i == 0 || (Abs(report[i]-report[i-1]) <= 3 &&
			report[i] != report[i - 1]) {
				continue
			} else {
				return false
			}
		}
	} else {
		return false
	}
	fmt.Println("This is safe:", report)
	return true
}

func isSafeReportWithout(report []int) bool {

	for i := range report {
		currArr := make([]int, 0, len(report)-1)
		currArr = append(currArr, report[:i]...)
		currArr = append(currArr, report[i+1:]...)
		if isSafeReports(currArr) {
			return true
		}
	}
	return false
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
