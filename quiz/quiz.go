package main
import (
		"flag"
		"fmt"
		"os"
	   )

func main()  {
	// Read command line flags
	csv_path := flag.String("csv", "questions.csv", "CSV with questions")
	flag.Parse()	
	file, e := os.Open(*csv_path)
	if e != nil {
		fmt.Println("Failed to open csv file")
		os.Exit(1)
	}
}
