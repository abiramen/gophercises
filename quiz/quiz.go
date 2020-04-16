package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Read command line flags
	csvPath := flag.String("csv", "questions.csv", "CSV with questions")
	flag.Parse()
	file, e := os.Open(*csvPath)
	if e != nil {
		errorExit(fmt.Sprintf("Failed to open csv file at %s", *csvPath))
	}
	r := csv.NewReader(file)
	pairs, e := r.ReadAll()
	if e != nil {
		fmt.Println("File could not be parsed.")
	}
	fmt.Println(pairs)
}

func errorExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
