package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"advent-of-code-2017/util"
	"fmt"
)

func main() {
	fmt.Println(computeEvenlyDivisorsSum("2.in"))
}

func computeEvenlyDivisorsSum(filename string) int {
	// read file
	file, _ := os.Open(filename)

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

	// for each row
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), "\t")

		// for each number in a row
		for i := 0; i < len(l); i++ {
			n, err := strconv.Atoi(l[i])
			util.Check(err)

			// for each next number in the same row
			for j:= i + 1; j < len(l); j++ {
				m, err := strconv.Atoi(l[j])
				util.Check(err)

				// check if evenly divides
				if n < m {
					if m % n == 0 {
						result += m / n
						break
					}
				} else {
					if n % m == 0 {
						result += n / m
						break
					}
				}
			}
		}
	}

	return result
}