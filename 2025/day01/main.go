package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	// "rsc.io/quote"
)

func main() {
	var final int
	value := 50
	// b := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	f, _ := os.Open("/home/machine/self/adventOfCode/2025/day01/input.txt")
	scanner := bufio.NewScanner(f)
	var number int
	for scanner.Scan() {
		number++
		input := scanner.Text()
		var cmd string = input[0:1]
		change, err := strconv.Atoi(input[1:])
		if err != nil {
			panic(err)
			return
		}
		change = change % 100
		switch cmd {
		case "L":
			if value-change < 0 {
				value = int(math.Abs(float64((change - value - 100) % 100)))
			} else if value-change == 100 {
				value = 0
			} else {
				value = value - change
			}

		case "R":
			if value+change > 100 {
				value = int(math.Abs(float64((value + change) % 100)))
			} else if value+change == 100 {
				value = 0
			} else {
				value = value + change
			}
		}

		if value == 0 {
			final++
		}
	}
	fmt.Println("part One", final, number)
}
