package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var partB = flag.Bool("partB", false, "Perform part B solution")

func main() {

	flag.Parse()

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var result int

	scanner := bufio.NewScanner(file)
	frequencies := make([]int, 0)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		frequencies = append(frequencies, i)
	}

	if !*partB {
		for _, i := range frequencies {
			result += i
		}
	} else {
		found := make(map[int]bool)
		done := false

		for !done {
			for i := 0; !done && i < len(frequencies); i++ {
				result += frequencies[i]
				if found[result] {
					done = true
				}
				found[result] = true
			}
		}
	}

	fmt.Printf("Result: %v\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
