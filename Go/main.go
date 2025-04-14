package main

import "fmt"

const currentDay = 2

func main() {

	sols := getSolutions()
	s1, p1, s2, p2 := sols[currentDay].callback()
	PrintSolutions(s1, p1, s2, p2)

}

func PrintSolutions(sample1, part1, sample2, part2 int) {
	fmt.Printf("Part 1 Sample Answer: %d\n", sample1)
	fmt.Printf("Part 1 Answer: %d\n", part1)
	fmt.Printf("Part 2 Sample Answer: %d\n", sample2)
	fmt.Printf("Part 2 Answer: %d\n", part2)
}

type solution struct {
	name     string
	callback func() (int, int, int, int)
}

func getSolutions() []solution {

	return []solution{
		{}, // blank to keep index aligned with day number
		{
			name:     "Day1",
			callback: Day1,
		},
		{
			name:     "Day2",
			callback: Day2,
		},
	}

}
