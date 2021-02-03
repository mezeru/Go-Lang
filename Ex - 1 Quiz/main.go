package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type prob struct {
	ques string
	ans  string
}

func parseLines(lines [][]string) []prob { // Creating a List from the data in CSV

	qp := make([]prob, len(lines))

	for i, line := range lines {
		qp[i] = prob{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}

	return qp
}

func main() {

	count := 0 // Counter for marks

	secs := 30 // Change to modify quiz time

	file, err := os.Open("C:/Users/Mezeru/Desktop/Go Prac/Ex - 1 Quiz/questions.csv")

	if err != nil {
		fmt.Print("Failed to open due to ", err, " \n")
		os.Exit(1)
	}

	reader := csv.NewReader(file) // Reader for CSV File

	lines, err := reader.ReadAll()
	list := parseLines(lines) // Parsing lines and creatring a list

	timer := time.NewTimer(time.Second * time.Duration(secs)) // Timer for quiz

	for _, no := range list {

		fmt.Print(no.ques, " = ")
		anschan := make(chan string)
		go func() {
			var userAns string
			fmt.Scanf("%s\n", &userAns) // Go Routine for answer
			anschan <- userAns
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime is UP !!!\n\nYour Score is ", count, " out of ", len(list))
			return

		case userAns := <-anschan:
			if userAns == no.ans {
				count = count + 1
				fmt.Println("Correct!!\n")
			} else {
				fmt.Println(" ")
			}

		}

	}

	fmt.Println("Your Score is ", count, " out of ", len(list))

}
