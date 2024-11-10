package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	entities "JG.Go.Quiz.Game/JG.entities"
	helper "JG.Go.Quiz.Game/JG.helper"
)

// =======================
// package level variables
const gameName string = "GO Quiz Game"

// var questions = make([]map[string]string, 0)
var userQuiz entities.UserData

// waits for the launched goroutine to finish
var wg = sync.WaitGroup{}

func main() {

	// welcome message to users
	greetUser()

	// populate the questions into quiz
	userQuiz.Questions = populateQuestions()

	userQuiz = getNumberOfQuestions()

	for {

		// get the user info
		userQuiz = getUserInput()

		var isValidName, isValidEmail = helper.ValidateUserInput(userQuiz)

		if !isValidName || !isValidEmail {
			if !isValidName {
				fmt.Println("Invalid name! Please input at least 2 chars.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email! Please input a valid email.")
			}
		} else {
			// start the quiz making questions
			userQuiz = makeQuestions()

			// show the total score at the end
			calculateTotalScore()

			// send the email to user
			go sendEmailToUser()

			break
		}
	}
}

func makeQuestionListRandom(questions []entities.Question) {
	rand.NewSource(time.Now().Unix())
	for i := len(questions) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}
}

func populateQuestions() []entities.Question {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Decode the JSON data
	var questions []entities.Question
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&questions)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	makeQuestionListRandom(questions)

	return questions
}

func greetUser() {
	fmt.Printf("Welcome to %v \n", gameName)
}

func getNumberOfQuestions() entities.UserData {
	var totalQuestionsToMake int
	var limitOfQuestionsToMake int = len(userQuiz.Questions)

	for {
		fmt.Printf("how many questions do you want to solve for this game? It is up to %v. \n", limitOfQuestionsToMake)
		fmt.Scan(&totalQuestionsToMake)
		if totalQuestionsToMake <= limitOfQuestionsToMake {
			break
		} else {
			fmt.Printf("Invalid answer! The limit of questions to answer is up to %v: \n", limitOfQuestionsToMake)
			continue
		}
	}

	userQuiz.TotalQuestionsToMake = totalQuestionsToMake

	return userQuiz
}

func getUserInput() entities.UserData {
	var name string
	var email string

	//ask users for their name & pointer
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	userQuiz.Name = name
	userQuiz.Email = email

	return userQuiz
}

func makeQuestions() entities.UserData {
	var userAnswer string
	var rightAnswer entities.Answer

	fmt.Println("Great! Let's start our quiz game: ")

	var index int = 1

	for _, question := range userQuiz.Questions {
		if userQuiz.TotalQuestionsToMake < index {
			break
		}

		fmt.Println(question.Title)

		for _, answer := range question.Answers {
			fmt.Println(answer.Title)
			if answer.IsRight {
				rightAnswer = answer
			}
		}

		fmt.Scan(&userAnswer)

		//if questions is correct store in the user info
		if userAnswer == rightAnswer.Code {
			fmt.Printf("Correct Answer! \n")
			fmt.Println("#################################")
			userQuiz.TotalScoreCorrect++
		} else {
			fmt.Printf("Wrong Answer! The correct answer is: %v \n", rightAnswer.Title)
			fmt.Println("#################################")
			userQuiz.TotalScoreWrong++
		}
		index++
	}
	return userQuiz
}

func calculateTotalScore() {
	fmt.Printf("Thank you for being part of our quiz game %v \n", userQuiz.Name)
	fmt.Println("Total Score")
	fmt.Printf("Right anwsers: %v \n", userQuiz.TotalScoreCorrect)
	fmt.Printf("Wrong anwsers: %v \n", userQuiz.TotalScoreWrong)
}

func sendEmailToUser() {
	//simulating 10s of delay
	time.Sleep(10 * time.Second)

	fmt.Println("#################################")
	fmt.Printf("Correct answers: %v \n", userQuiz.TotalScoreCorrect)
	fmt.Printf("Wrong answers: %v \n", userQuiz.TotalScoreWrong)
	fmt.Println("#################################")

	// auth := smtp.PlainAuth(
	// 	"",
	// 	"email@gmail.com",
	// 	"",
	// 	"smpt.gmail.com",
	// )

	// smpt.SendMail

	wg.Done()
}
