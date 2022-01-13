package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "please provide a csv file in questions|answers format")
	limit := flag.Int("limit", 30, "please provide the limit limit in seconds")
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

	problems := readLines(lines)

	timer := time.NewTimer(time.Second * time.Duration(*limit))

	var score uint
problemsLoop:
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)

		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case answer := <-answerChan:
			if answer == p.a {
				score++
			}
		case <-timer.C:
			break problemsLoop
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
