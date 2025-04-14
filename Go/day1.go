package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() (int, int, int, int) {

	const samplePath = "../Data/Sample_Day1.txt"
	const inputPath = "../Data/Day1.txt"

	sFile, err := os.Open(samplePath)
	if err != nil {
		log.Fatalf("could not open %s: %s", inputPath, err)
	}
	defer sFile.Close()

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("could not open %s: %s", inputPath, err)
	}
	defer file.Close()

	l1, l2 := D1.CreateLists(sFile)
	s1 := D1.part1Calc(l1, l2)
	s2 := D1.part2Calc(l1, l2)

	l1, l2 = D1.CreateLists(file)
	p1 := D1.part1Calc(l1, l2)
	p2 := D1.part2Calc(l1, l2)

	return s1, p1, s2, p2

}

var D1 = struct {
	CreateLists func(*os.File) ([]int, []int)
	part1Calc   func([]int, []int) int
	part2Calc   func([]int, []int) int
}{

	CreateLists: func(f *os.File) ([]int, []int) {

		scanner := bufio.NewScanner(f)
		var l1, l2 []int

		for scanner.Scan() {
			cols := strings.SplitN(scanner.Text(), "   ", 2)

			i, _ := strconv.Atoi(cols[0])
			l1 = append(l1, i)
			i, _ = strconv.Atoi(cols[1])
			l2 = append(l2, i)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return l1, l2

	},

	// Calculate total distances between list
	part1Calc: func(l1, l2 []int) int {

		slices.Sort(l1)
		slices.Sort(l2)
		var dist float64

		for i, val := range l1 {
			dist += math.Abs(float64(val - l2[i]))
		}

		return int(dist)

	},

	// Calculate similarity score
	part2Calc: func(l1, l2 []int) int {

		counts := make(map[int]int)
		sim := 0

		for _, val := range l1 {
			counts[val] = 0
		}

		for _, val := range l2 {
			if _, ok := counts[val]; ok {
				counts[val] += 1
			}
		}

		for _, val := range l1 {
			sim += val * counts[val]
		}

		return sim

	},
}
