package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

func loadQuestions(filename string) ([]Question, error) {
	var questions []Question
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}

func main() {
	questions, err := loadQuestions("quiz.json")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		os.Exit(1)
	}

	score := 0
	for _, q := range questions {
		fmt.Println(q.Question)
		for i, option := range q.Options {
			fmt.Printf("%d: %s\n", i+1, option)
		}

		var answer int
		fmt.Print("Your answer (1-4): ")
		fmt.Scan(&answer)

		if answer-1 == q.Answer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}
		fmt.Println()
	}

	fmt.Printf("Your final score: %d out of %d\n", score, len(questions))
}
