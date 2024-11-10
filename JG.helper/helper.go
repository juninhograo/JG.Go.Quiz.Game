package helper

import (
	"strings"

	entities "JG.Go.Quiz.Game/JG.entities"
)

func ValidateUserInput(userData entities.UserData) (bool, bool) {
	isValidName := len(userData.Name) >= 2
	isValidEmail := strings.Contains(userData.Email, "@") && len(userData.Email) >= 3
	return isValidName, isValidEmail
}
