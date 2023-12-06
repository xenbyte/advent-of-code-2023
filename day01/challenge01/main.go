package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: go run main.go <input_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := getFirstAndLastDigits(line)
		result := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		resultNum, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		count = count + resultNum
	}

	fmt.Println(count)
}

func getFirstAndLastDigits(s string) (int, int) {
	runes := []rune(s)

	var firstDigit int
	for _, r := range runes {
		if r >= '0' && r <= '9' {
			firstDigit = int(r - '0')
			break
		}
	}

	var lastDigit int
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			lastDigit = int(runes[i] - '0')
			break
		}
	}

	return firstDigit, lastDigit
}
