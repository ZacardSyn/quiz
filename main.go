package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "problems in csv format")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		errorMessage(fmt.Sprintf("Failed to open file: %s, please make sure it's a csv file\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		errorMessage("Can't read csv file")
	}
	problemStruct := parseLines(lines)
	correct := float32(0)
	for i, problem := range problemStruct {
		fmt.Printf("Problem %d: %s\n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("Correct!")
			correct++
		}else{
			fmt.Println("Incorrect!")
		}
	}
	percentage := correct / float32(len(lines)) * 100
	fmt.Printf("You got %d questions right, you score is %.2f%\n", correct, percentage)
}

func errorMessage(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	retProblem := make([]problem, len(lines))
	for i, line := range (lines){
		retProblem[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return retProblem
}