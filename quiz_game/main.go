package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

//make a struct to help specify the contents of the csv file
type pro struct {
	question string
	answer   string
}

func main() {
	//define flags to read command line arguments input
	//filename specifies the filename
	filename := flag.String("csv", "problems.csv", "This file program expects a csv input")
	//limit specifies the game duration
	limit := flag.Int("limit", 4, "The time in seconds it takes to complete the exercise")
	//add a flag to either shuffle the slice or not
	shuffle := flag.Bool("shuffle", false, "Change to true or false to shuffe the questions")
	flag.Parse()

	//open the csv and check for error
	problem, err := os.Open(*filename)
	if err != nil {
		fmt.Println("An error encountered: File doesn't exist. Filename you entered is", *filename)
	} else {
		//get the newreader method and read the file, save the output into a variable
		file := csv.NewReader(problem)
		output, err := file.ReadAll()
		if err == io.EOF {
			fmt.Println("This is the end of the file")
		}

		//shuffle the csv file
		if *shuffle == true {
			rand.Seed(time.Now().Unix())
			rand.Shuffle(len(output), func(i, j int) {
				output[i], output[j] = output[j], output[i]
			})
		}
		//use the timer function to set liit of the duration. It gets the limit from the flas or use the default
		timer := time.NewTimer(time.Duration(*limit) * time.Second)
		//make a chan that will receive the answers
		ans := make(chan string)
		//save the scored into the score variale which has initial value of 0
		score := 0
		//loop through the read file and get the individual colunms
		for numb, eachLine := range output {
			qAndA := pro{
				question: strings.TrimSpace(eachLine[0]),
				answer:   strings.TrimSpace(eachLine[1]),
			}
			fmt.Printf("Question %v is %v =\n", numb+1, qAndA.question)
			//use goroutine to accept input from user so that it won't block the limit channel from doing its task
			go func() {
				var inp string
				fmt.Scanln(&inp)
				ans <- inp
			}()

			select {
			//return and exit when the time reaches and print the total score
			case <-timer.C:
				fmt.Printf("\nYou scored total of %v out of %v\n", score, len(output))
				return
			case correct := <-ans:
				//receive answer from the channel and compare it with the answer in the csv file
				if qAndA.answer == strings.TrimSpace(correct) {
					fmt.Println("your answer is right.")
					//increament the score
					score++
				} else {
					fmt.Println("Answer is wrong")
				}
			}
		}
		//print the total score if all questions are succesfully answered
		fmt.Printf("You scored total of %v out of %v\n", score, len(output))

	}
	defer problem.Close()
}
