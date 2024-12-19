package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int64
	// var resultDo int64
	strSlice := make([]string, 0)
	for scanner.Scan() {

		scannedText := scanner.Text()
		strSlice = append(strSlice, scannedText)
		
		result += findAllMuls(scannedText)
		
		// resultDo += findDoDontMuls(scannedText)
		// fmt.Println("Scanned text:", scannedText)
		
		// fmt.Println("Found:", result)
	}
	fmt.Println("Result is:", result)
	
	joinedStr := strings.Join(strSlice, "")
	// fmt.Println("joinedStr is:", strSlice)
	resultDo := findDoDontMuls(joinedStr)
	fmt.Println("Result of muls between do/don't is:", resultDo)
}

func findAllMuls(scannedText string) int64 {
	var result int64
	mulInput := regexp.MustCompile(`mul\(\d+,\d+\)`)

	res := mulInput.FindAllString(scannedText, -1)
	for _, resString := range res {
		pair := regexp.MustCompile(`\d+`).FindAllString(resString, -1)
		// fmt.Println("This is A:", pair)
		a, err := strconv.Atoi(pair[0])
		if err != nil {
			fmt.Println("Convert to a is not successful:", err)
		}

		b, err := strconv.Atoi(pair[1])
		if err != nil {
			fmt.Println("Convert to b is not successful:", err)
		}
		result += int64(a) * int64(b)
	}
	return result
}

func findDoDontMuls(scannedText string) int64 {
	var result int64
	doDontMul := regexp.MustCompile(`don't\(\)(.*?)do\(\)`)
	resStr := doDontMul.ReplaceAllString(scannedText, "")
	fmt.Println(resStr)
	dontStrs := regexp.MustCompile(`don't\(\).*`)
	finStr := dontStrs.ReplaceAllString(resStr, "")
	fmt.Println("final", finStr)
	result += findAllMuls(finStr)
	return result
}
