package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var redLimit = 12
var greenLimit = 13
var blueLimit = 14

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

	validGameIDsSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		Calculate(line, &validGameIDsSum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	fmt.Printf("Sum of valid game IDs: %d\n", validGameIDsSum)

}

func Calculate(line string, validGameIDsSum *int) {
	if strings.HasPrefix(line, "Game ") {
		parts := strings.Split(line, ": ")
		gameIDPart := parts[0]
		cubesDataPart := parts[1]

		gameIDStr := strings.Split(gameIDPart, " ")[1]
		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			log.Fatalf("Invalid game ID: %v", err)
		}

		// Parse cube data
		validGame := true
		cubeSets := strings.Split(cubesDataPart, "; ")
		for _, cubeSet := range cubeSets {
			cubes := strings.Split(cubeSet, ", ")
			redCount := 0
			greenCount := 0
			blueCount := 0
			for _, cube := range cubes {
				cubeDetails := strings.Split(cube, " ")
				count, err := strconv.Atoi(cubeDetails[0])
				if err != nil {
					log.Fatalf("Invalid cube count: %v", err)
				}
				color := cubeDetails[1]
				switch color {
				case "red":
					redCount += count
				case "green":
					greenCount += count
				case "blue":
					blueCount += count
				}
			}
			if redCount > redLimit || greenCount > greenLimit || blueCount > blueLimit {
				validGame = false
				break
			}
		}

		if validGame {
			*validGameIDsSum += gameID
		}
	}
}
