package main

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day2() (int, int, int, int) {

	const samplePath = "../Data/Sample_Day2.txt"
	const inputPath = "../Data/Day2.txt"

	sFile, err := os.Open(samplePath)
	if err != nil {
		log.Fatalf("could not open %s: %s", samplePath, err)
	}
	defer sFile.Close()

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("could not open %s: %s", inputPath, err)
	}
	defer file.Close()

	sReports := D2.ParseReports(sFile)
	s1 := D2.part1Calc(sReports)
	s2 := D2.part2Calc(sReports)

	reports := D2.ParseReports(file)
	p1 := D2.part1Calc(reports)
	p2 := D2.part2Calc(reports)

	return s1, p1, s2, p2
}

type report struct {
	levels []int
}

func (rpt *report) IsSafe() bool {

	delta := make([]int, len(rpt.levels)-1)
	for i := 1; i < len(rpt.levels); i++ {
		delta[i-1] = rpt.levels[i] - rpt.levels[i-1]
	}
	// fmt.Printf("Report %v ==> Delta %v\n", *rpt, delta)

	// determine if slope is all positive or negative
	if delta[0] == 0 {
		return false
	}

	lastSign := cmp.Compare(delta[0], 0)
	for i := 1; i < len(delta); i++ {
		if delta[i] == 0 {
			return false
		}

		if cmp.Compare(delta[i], 0) != lastSign {
			return false
		}
	}

	return (slices.Max(delta) <= 3) && (slices.Min(delta) >= -3)
}

func (rpt *report) IsDampenSafe() bool {

	if rpt.IsSafe() {
		return true
	}

	for i, _ := range rpt.levels {

		modR := report{levels: make([]int, len(rpt.levels))}
		copy(modR.levels, rpt.levels)
		modR.levels = slices.Delete(modR.levels, i, i+1)

		if modR.IsSafe() {
			return true
		}
	}

	return false
}

var D2 = struct {
	ParseReports func(*os.File) []report
	part1Calc    func([]report) int
	part2Calc    func([]report) int
}{

	ParseReports: func(f *os.File) []report {

		scanner := bufio.NewScanner(f)

		rptSet := []report{}

		for i := 0; scanner.Scan(); i++ {

			strVals := strings.Split(scanner.Text(), " ")
			rpt := report{levels: make([]int, len(strVals))}

			for j, entry := range strVals {
				// assume data input is clean and no errors
				rpt.levels[j], _ = strconv.Atoi(entry)
			}

			rptSet = append(rptSet, rpt)

		}

		return rptSet
	},

	part1Calc: func(rpts []report) int {

		sum := 0
		for _, rpt := range rpts {
			if rpt.IsSafe() {
				sum++
			}
		}

		return sum
	},

	part2Calc: func(rpts []report) int {

		sum := 0
		for _, rpt := range rpts {
			if rpt.IsDampenSafe() {
				sum++
			}
		}

		return sum
	},
}
