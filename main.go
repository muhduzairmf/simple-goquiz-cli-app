package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type user struct {
	name string
	age int
	totalMark int
}

type question struct {
	statement string
	answer string 
}

func getInput(prompt string, inputReader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	userInput, err := inputReader.ReadString('\n')
	
	return strings.TrimSpace(userInput), err
}

func greeting(userName *string) {
	fmt.Println("Welcome to quiz game.")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter your name : ", inputReader)

	fmt.Println("\nHello " + name + ", welcome to the game!")
	// Or,
	// fmt.Printf("\nHello %v, welcome to the game!", name)

	*userName = name
}

func isUserQualified(userAge *int) string {
	
	inputReader := bufio.NewReader(os.Stdin)
	age_input, _ := getInput("Enter your age : ", inputReader)

	age, err := strconv.ParseInt(age_input, 10, 0)

	if err != nil {
		fmt.Println("Age must be a valid number")
		return "Invalid input"
	} 

	if age < 10 {
		fmt.Println("\nYou are not old enough to play. Bye....")
		return "Too young"
	}

	fmt.Println("\nYou can play this game. Let's go....")
	*userAge = int(age)

	return "Qualified"
}

func (theQuestion *question) askQuestion(currentMark *int) int {
	
	inputReader := bufio.NewReader(os.Stdin)
	userAnswer, _ := getInput(theQuestion.statement + "\n>> ", inputReader)

	if strings.ToLower(userAnswer) != strings.ToLower(theQuestion.answer) {
		fmt.Println("Incorrect!")
	} else {
		fmt.Println("Correct!")
		*currentMark++
	}

	return *currentMark
}

func announceResult(theUser *user) {
	fmt.Println("\nThank you for playing, " + theUser.name)
	fmt.Println("Total marks " + strconv.Itoa(theUser.totalMark) + "/3")
	// strconv.Itoa() function is convert int to string
}

func main()  {

	var newUser user
	
	greeting(&newUser.name)

	var ageStatus string

	for ageStatus != "Qualified" {
		ageStatus = isUserQualified(&newUser.age)

		if ageStatus == "Too young" {
			fmt.Print("Press enter to exit....")
			fmt.Scanln()
			return
		}
	}

	newUser.totalMark = 0

	var question1 question

	question1.statement = "What is file extension for Go Programming Language? (.go/.g/.golang)"
	question1.answer = ".go"

	question1.askQuestion(&newUser.totalMark)

	var question2 question

	question2.statement = "Who is design the Go Programming Language? (JetBrains/Google/Microsoft)"
	question2.answer = "Google"

	question2.askQuestion(&newUser.totalMark)

	var question3 question

	question3.statement = "What is the official Go website? (golang.org/go.io/go.dev)"
	question3.answer = "go.dev"

	question3.askQuestion(&newUser.totalMark)

	announceResult(&newUser)
}
