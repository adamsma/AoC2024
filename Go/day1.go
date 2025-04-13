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

const samplePath = "../Data/Sample_Day1.txt"
const inputPath = "../Data/Day1.txt"

func Day1() (int, int, int, int) {

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

	l1, l2 := CreateLists(sFile)
	s1 := part1Calc(l1, l2)
	s2 := part2Calc(l1, l2)

	l1, l2 = CreateLists(file)
	p1 := part1Calc(l1, l2)
	p2 := part2Calc(l1, l2)

	return s1, p1, s2, p2

}

func CreateLists(f *os.File) ([]int, []int) {

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

}

// Calculate total distances between list
func part1Calc(l1, l2 []int) int {

	slices.Sort(l1)
	slices.Sort(l2)
	var dist float64

	for i, val := range l1 {
		dist += math.Abs(float64(val - l2[i]))
	}

	return int(dist)

}

// Calculate similarity score
func part2Calc(l1, l2 []int) int {

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

}
