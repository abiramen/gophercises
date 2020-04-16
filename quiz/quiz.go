package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Store question-answer pairs
type quiz struct {
	question string
	answer   string
}

func main() {
	// Read a command line flag specifying the CSV with the questions.
	csvPath := flag.String("csv", "questions.csv", "CSV with questions")
	flag.Parse()

	// Open the file from path.
	file, e := os.Open(*csvPath)
	if e != nil {
		// Handle any errors.
		errorExit(fmt.Sprintf("Failed to open csv file at %s", *csvPath))
	}

	// Read the file as a CSV.
	r := csv.NewReader(file)
	pairs, e := r.ReadAll()
	if e != nil {
		// Handle any errors.
		errorExit("File could not be parsed.")
	}
	// Process the resulting strings.
	quiz := parsePairs(pairs)

	score := 0
	// Iterate through each question.
	for index, problem := range quiz {
		fmt.Printf("Problem #%d: %s = ", index+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}
	}
	fmt.Printf("You scored %d out of %d correct!\n", score, len(quiz))

}

// Convert strings taken from the CSV to quiz problems.
func parsePairs(pairs [][]string) []quiz {
	// Allocates space for problems depending on the number of string pairs.
	ret := make([]quiz, len(pairs))
	for i, pair := range pairs {
		ret[i] = quiz{
			question: pair[0],
			answer:   strings.TrimSpace(pair[1]),
		}
	}
	return ret
}

func errorExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
