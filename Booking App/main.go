// Firs thing to start a project in golang is open the terminal and use de followed sentence 'go mod init <BookingApp>.'
package main // In go, we organize our project in packages, and this is the main.

import (
	"fmt" // It's the 'format' package, it makes a lot of functions available.
	"strings"
)

func main() { // The main function is where the code starts to run, without it, the programme won't know were to start.
	// go run 'main.go' is the command to run a programme in go.

	var conferenceName = "Go Conference" // It's declared as a variable because if we need to change the name, it won't necessary to change all the code.
	const conferenceTickets = 50         // It's declared as a constant because it's value won't change and if we try to change the value, it won't work.
	remainingTickets := 50               // These two declarations above are called inference declaration and this one is called short declaration and can't be used when declaring a constant.
	//var bookings [50]string  			// The arrays in go have fixed sizes, so it's important to inform the maximum size this array can hold, when declaring. We are able to add and exclude elements of the array during execution.
	bookings := []string{} // The slices in go are flexible, and has a variable lenght because it's not necessary inform it's size at the beginning,

	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets %T\n", conferenceName, conferenceTickets, remainingTickets)  // The %T can be used to see the value of a variable.

	fmt.Printf("Welcome to %v booking application!\n", conferenceName) // It's the formated print, makes easier to writte things as below.
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	for { // It's used as a infinit loop.
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// fmt.Println(remainingTickets)  Here we can see the value stored in the variable (50).
		// fmt.Println(&remainingTickets)  And with this pointer indicated by &, we are able to see the address in memory of this variable (0xc00000a0b8)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 // Here we are checking some conditions to see if the names are valid.
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//isValidCity := city == "Singapore" || city == "London"
		// !isValidCity  By doing that, we negate the variable.

		if isValidName && isValidEmail && isValidTicketNumber { // It's a simple if/else statment.
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName  This is how we would do if we had to work with an array, and as we don´t know the exactly number of people that are interested in the tickets and the first empty position to add this person, it's easier to use a slice.
			bookings = append(bookings, firstName+" "+lastName) // This is how we add someone to the slice.

			// Informations about the array.
			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array lenght: %v\n", len(bookings))

			// Informations about the slice.
			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice lenght: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // This '_' is used to ignore a varible that I don't want to use.
				var names = strings.Fields(booking) // This function splits the string with white spaces as separator and returns a slice with the split elements ["Guilherme","Duran"].
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 { // It's a simple if statment.
				fmt.Println("Our conference is booked out. Come back next year.")
				break // It's used as a stop order.
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email adress you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered is invalid")
			}
			//continue  // It goes to next iteration
		}
	}
}
