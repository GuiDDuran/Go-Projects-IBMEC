package helper // Normally when we are organizing our code, we put all the files that belongs to a package, in the same directory.
// As the package is defined as main, it can identify the variables from the main file and to run, We have to iclude all the files. 'go run main.go helper.go'or 'go run .'

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) { // When I more than one return, I have to especify all the types of it returns and by using up-letter in the name's function, we make it available to be used in the main file.
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 // Here we are checking some conditions to see if the names are valid.
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
