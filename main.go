package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	argsWithoutProg := os.Args[1:]

	var err error
	nums := make([]int, len(argsWithoutProg))
	for i := 0; i < len(argsWithoutProg); i++ {
		if nums[i], err = strconv.Atoi(argsWithoutProg[i]); err != nil {
			panic(err)
		}
	}

	rate := nums[0]
	total := nums[1]

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
