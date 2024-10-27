package main

import "fmt"

func midCal() int {
	var midScore int
	fmt.Print("Input your mid term scores ")
	fmt.Scanf("%d", &midScore)
	for midScore > 60 || midScore < 0 {
		fmt.Println("Your score is incorrect. Please input your score again")
		fmt.Print("Input your mid term scores ")
		fmt.Scanf("%d", &midScore)
	}
	return midScore
}

func finalCal() int {
	var finalScore int
	fmt.Print("Input your final term scores ")
	fmt.Scanf("%d", &finalScore)
	for finalScore > 40 || finalScore < 0 {
		fmt.Println("Your score is incorrect. Please input your score again")
		fmt.Print("Input your final term scores ")
		fmt.Scanf("%d", &finalScore)
	}
	return finalScore
}

func specialScore() int {
	var specialScore int
	fmt.Print("Input your special scores ")
	fmt.Scanf("%d", &specialScore)
	for specialScore > 10 || specialScore < 0 {
		fmt.Println("Your score is incorrect. Please input your score again")
		fmt.Print("Input your special scores ")
		fmt.Scanf("%d", &specialScore)
	}
	return specialScore
}

func calScores() int {
	midScore := midCal()
	finalScore := finalCal()
	specialScore := specialScore()
	return midScore + finalScore + specialScore
}

func calGrade(score int) {
	if score > 100 || score < 0 {
		fmt.Println("Your score is incorrect")
	} else if score > 80 {
		fmt.Println("Your grade is A")
	} else if score > 70 {
		fmt.Println("Your grade is B")
	} else if score > 60 {
		fmt.Println("Your grade is C")
	} else if score > 50 {
		fmt.Println("Your grade is D")
	} else {
		fmt.Println("Your grade is F")
	}
}

func main() {
	score := calScores()
	calGrade(score)
}
