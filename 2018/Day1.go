package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input/Day1.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")

	FreqDict := map[int]int{}
	DuplicateNotFound := true
	CurrentFreq := 0
	idx := 0
	i := 0
	// for i := 0; i < len(lines); i++{
	for DuplicateNotFound {
		if len(lines[idx]) == 0 {
			i += 1
			idx = (idx + 1) % len(lines)
			continue
		}

		FreqDict[CurrentFreq] = FreqDict[CurrentFreq] + 1
		if FreqDict[CurrentFreq] == 2 {
			break
		}

		ChangeDir := lines[idx][0]
		ChangeFreq, err := strconv.Atoi(lines[idx][1:])
		if err != nil {
			panic(err)
		}
		
		NewFreq := 0
		switch ChangeDir {
		case '+':
			NewFreq = CurrentFreq + ChangeFreq
		case '-':
			NewFreq = CurrentFreq - ChangeFreq
		default:
			panic("Unknown Direction: " + string(ChangeDir))
		}

		fmt.Printf("Loop: %d, Freq %d; Current frequency %d, change of %c%d; resulting freqency %d.\n",
			(i / len(lines)), i%len(lines), CurrentFreq, ChangeDir, ChangeFreq, NewFreq,
		)

		CurrentFreq = NewFreq
		i += 1
		idx = (idx + 1) % len(lines)
	}

	fmt.Printf("Duplicated Frequency Found: %d.", CurrentFreq)
}
