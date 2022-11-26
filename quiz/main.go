package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem Struct for holding question and answers
type problem struct {
	Q, A string
}

func main() {
	// CSV flag for collecting problem file
	csvFileName := flag.String("csv", "problems.csv", "A CSV file in  the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "The limit for the quiz in seconds")
	flag.Parse()

	// Open file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	// New reading receiving an io.Reader (in this case a file as it satisfies and implements the io.Reader interface)
	r := csv.NewReader(file)
	// Read all contents of the file into a slice,
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV")
	}

	// Parse lines(csv content) to a slice  of problems
	problems := parseLines(lines)
	timmer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// <-timmer.C
	// correct answer counter
	correct := 0
	answerChan := make(chan string)
	// loop through the problems
	fmt.Printf("%s %d %s %d %s \n\n", "You have only ", *timeLimit, "second(s) to complete ", len(problems), " questions.")
	for i, problem := range problems {
		fmt.Printf("Problem #%d, %s =\n", i+1, problem.Q)
		//Create go routines for
		go func() {
			var answer string
			// Get answer from standard input
			fmt.Scanf("%s", &answer)
			answerChan <- answer
		}()
		select {
		case <-timmer.C:
			fmt.Printf("\n %s \n", "You have exceeded your alloted time.")
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == problem.A {
				correct++
			}

		}

	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, 0)
	for _, line := range lines {
		problems = append(problems, problem{
			line[0], strings.TrimSpace(line[1]),
		})
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
