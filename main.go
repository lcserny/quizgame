package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "please provide a csv file in questions|answers format")
	flag.Parse()

	csvFile, err := os.Open(*csvPath)
	if err != nil {
		exit(fmt.Sprintf("could not read file provided %s", *csvPath))
	}

	csvReader := csv.NewReader(csvFile)
	lines, err := csvReader.ReadAll()
	if err != nil {
		exit("csv file could not be parsed")
	}

	var score uint
	problems := readLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, len(problems))
}

func readLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
