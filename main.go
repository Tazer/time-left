package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	argsWithoutProg := os.Args[1:]

	var err error

	rate, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	total, err := convertShorthandToNumber(argsWithoutProg[1])
	if err != nil {
		panic(err)
	}

	f, d := timeLeft(rate, total)

	fmt.Printf("With a rate of %d, and a total of %d, the time left is:\n", rate, total)
	fmt.Println(d)
	fmt.Println("The finish time is:")
	fmt.Println(f.Format(time.RFC3339))

}

func timeLeft(rate int, total int) (time.Time, time.Duration) {
	r := total / rate
	finishTime := (time.Now().Add(time.Duration(r) * time.Second))
	return finishTime, finishTime.Sub(time.Now())
}

// convertShorthandToNumber converts shorthand notations K, M, and B to their numeric values.
func convertShorthandToNumber(input string) (int, error) {
	// Remove spaces and make the input uppercase for standardization
	trimmedInput := strings.ToUpper(strings.TrimSpace(input))
	if len(trimmedInput) == 0 {
		return 0, errors.New("input is empty")
	}

	// Check the last character for K, M, or B and multiply accordingly
	lastChar := trimmedInput[len(trimmedInput)-1:]
	var multiplier int

	switch strings.ToUpper(lastChar) {
	case "K":
		multiplier = 1_000
	case "M":
		multiplier = 1_000_000
	case "B":
		multiplier = 1_000_000_000
	default:
		// If there's no K, M, or B, try to convert the input directly
		return strconv.Atoi(trimmedInput)
	}

	// Parse the number before the K, M, or B
	number, err := strconv.Atoi(trimmedInput[:len(trimmedInput)-1])
	if err != nil {
		return 0, err
	}

	return number * multiplier, nil
}
