package entities

type UserData struct {
	Name                 string
	Email                string
	TotalScoreCorrect    int
	TotalScoreWrong      int
	TotalQuestionsToMake int
	Questions            []Question
}

type Question struct {
	Title   string
	Answers []Answer
}

type Answer struct {
	Title   string
	Code    string
	IsRight bool
}
