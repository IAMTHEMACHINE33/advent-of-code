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
	var partOne int
	var partTwo int
	value := 50
	// b := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	f, _ := os.Open("/home/machine/self/adventOfCode/2025/day01/input.txt")
	scanner := bufio.NewScanner(f)
	var number int
	// for i := range len(b) {
	for scanner.Scan() {
		number++
		input := scanner.Text()
		// input := b[i]
		var cmd string = input[0:1]
		change, err := strconv.Atoi(input[1:])
		if err != nil {
			panic(err)
			return
		}

		partTwo += change / 100
		change = change % 100
		// fmt.Println(cmd, value, change)
		switch cmd {
		case "L":
			if value-change < 0 {
				// fmt.Println("LIA", value-change)
				if value != 0 {
					partTwo++
				}
				value = int(math.Abs(float64((change - value - 100) % 100)))
				// fmt.Println("LI", partOne)
			} else if value-change == 100 {
				value = 0
			} else {
				value = value - change
			}

		case "R":
			if value+change > 100 {
				// fmt.Println("RIA", value-change)
				if value != 0 {
					partTwo++
				}
				value = int(math.Abs(float64((value + change) % 100)))
				// fmt.Println("RI", partOne)
			} else if value+change == 100 {
				value = 0
			} else {
				value = value + change
			}
		}

		if value == 0 {
			partOne++
		}
	}
	fmt.Println("part One", partOne)
	fmt.Println("part Two", partOne+partTwo)
}
