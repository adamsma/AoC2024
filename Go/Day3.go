package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day3() (int, int, int, int) {

	const samplePath = "../Data/Sample_Day3.txt"
	const samplePathP2 = "../Data/Sample_Day3p2.txt"
	const inputPath = "../Data/Day3.txt"

	sFile, err := os.Open(samplePath)
	if err != nil {
		log.Fatalf("could not open %s: %s", samplePath, err)
	}
	defer sFile.Close()

	s2File, err := os.Open(samplePathP2)
	if err != nil {
		log.Fatalf("could not open %s: %s", samplePathP2, err)
	}
	defer s2File.Close()

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("could not open %s: %s", inputPath, err)
	}
	defer file.Close()

	file2, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("could not open %s: %s", inputPath, err)
	}
	defer file2.Close()

	sData := D3.ParseCode(sFile)
	s1 := D3.part1Calc(sData)

	s2Data := D3.ParseCodeV2(s2File)
	s2 := D3.part2Calc(s2Data)

	data := D3.ParseCode(file)
	p1 := D3.part1Calc(data)

	data2 := D3.ParseCodeV2(file2)
	p2 := D3.part2Calc(data2)

	return s1, p1, s2, p2
}

const D3_MUL_REGEX = `mul\(\d+,\d+\)`

type MulFx struct {
	x0 int
	x1 int
}

func NewMulFx(instruction string) MulFx {

	ok, err := regexp.MatchString(D3_MUL_REGEX, instruction)
	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		log.Fatal(fmt.Errorf("malformed instruction text: %s", instruction))
	}

	re := regexp.MustCompile(`\d+`)
	s := re.FindAllString(instruction, 2)
	x0, _ := strconv.Atoi(s[0])
	x1, _ := strconv.Atoi(s[1])

	return MulFx{x0: x0, x1: x1}

}

func (m MulFx) Exec() int {
	return m.x0 * m.x1
}

var D3 = struct {
	ParseCode   func(*os.File) []MulFx
	ParseCodeV2 func(*os.File) []MulFx
	ExecMul     func(string) int
	part1Calc   func([]MulFx) int
	part2Calc   func([]MulFx) int
}{

	ParseCode: func(f *os.File) []MulFx {

		scanner := bufio.NewScanner(f)
		txt := ""

		for scanner.Scan() {
			txt += scanner.Text()
		}

		re := regexp.MustCompile(D3_MUL_REGEX)

		matches := re.FindAllString(txt, -1)
		instructions := make([]MulFx, len(matches))

		for i, fx := range matches {
			instructions[i] = NewMulFx(fx)
		}

		return instructions

	},

	ParseCodeV2: func(f *os.File) []MulFx {

		scanner := bufio.NewScanner(f)
		txt := ""

		for scanner.Scan() {
			txt += scanner.Text()
		}

		// deterimine when start of MUL_REGEX falls between start
		// of do regex and start of dont regex
		re := regexp.MustCompile(D3_MUL_REGEX)
		doRe := regexp.MustCompile(`do\(\)`)
		dontRe := regexp.MustCompile(`don't\(\)`)

		matches := re.FindAllString(txt, -1)
		mulLoc := re.FindAllStringIndex(txt, -1)
		doLoc := doRe.FindAllStringIndex(txt, -1)
		dontLoc := dontRe.FindAllStringIndex(txt, -1)

		// add implicit do at start and dont at end
		condLoc := slices.Concat(
			[][]int{{0, 0}}, doLoc, dontLoc, [][]int{{len(txt), len(txt)}},
		)
		condType := strings.Split(
			"do|"+strings.Repeat("do|", len(doLoc))+
				strings.Repeat("dont|", len(dontLoc))+"dont",
			"|",
		)

		// create type to all for ease of co-sorting
		type Loc struct {
			Index int
			Kind  string
		}

		locations := make([]Loc, len(condLoc))
		for i, loc := range condLoc {
			locations[i] = Loc{Index: loc[0], Kind: condType[i]}
		}

		slices.SortFunc(locations, func(a, b Loc) int {
			return cmp.Compare(a.Index, b.Index)
		})

		fmt.Println("locations: ", locations)

		instructions := make([]MulFx, 0)
		for i, loc := range mulLoc {
			pos, _ := slices.BinarySearchFunc(
				locations,
				Loc{loc[0], "mul"},
				func(a, b Loc) int { return cmp.Compare(a.Index, b.Index) },
			)

			if locations[pos-1].Kind == "do" {
				instructions = append(instructions, NewMulFx(matches[i]))
			}
		}

		return instructions

	},

	part1Calc: func(instructions []MulFx) int {

		sum := 0
		for _, mul := range instructions {
			sum += mul.Exec()
		}

		return sum

	},

	part2Calc: func(instructions []MulFx) int {

		sum := 0
		for _, mul := range instructions {
			sum += mul.Exec()
		}

		return sum

	},
}
