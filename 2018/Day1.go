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

	CurrentFreq := 0
	FreqMap := map[int]int{};
	FreqMap[CurrentFreq] += 1
	idx := 0;
	for {
		if len(lines[idx]) == 0 {
			idx = (idx + 1) % len(lines)
			continue
		}

		FreqChange, err := strconv.Atoi(lines[idx])
		if err != nil {
			panic(err)
		}

		NewFreq := CurrentFreq + FreqChange

		fmt.Printf("Current frequency %d, change of %d; resulting freqency %d.\n",
			CurrentFreq, FreqChange, NewFreq,
		)

		CurrentFreq = NewFreq
		FreqMap[CurrentFreq] += 1
		if (FreqMap[CurrentFreq] > 1){
			break
		}


		idx = (idx + 1) % len(lines)
	}

	fmt.Printf("Final Frequency Found: %d.", CurrentFreq)
}
