// Firs thing to start a project in golang is open the terminal and use de followed sentence 'go mod init <BookingApp>.'
package main // In go, we organize our project in packages, and this is the main.

import (
	"BookingApp/helper" // To use the helper package, we also have to put the module name in front of the name's package.
	"fmt"               // It's the 'format' package, it makes a lot of functions available.
	"sync"
	"time"
	//"strconv"
)

const conferenceTickets = 50         // It's declared as a constant because it's value won't change and if we try to change the value, it won't work.
var conferenceName = "Go Conference" // It's declared as a variable because if we need to change the name, it won't necessary to change all the code.
var remainingTickets = 50            // These two declarations above, with the ':=' are called inference declaration and this one is called short declaration and can't be used when declaring a constant.
// var bookings [50]string  			// The arrays in go have fixed sizes, so it's important to inform the maximum size this array can hold, when declaring. We are able to add and exclude elements of the array during execution.
var bookings = make([]UserData, 0) // The slices in go are flexible, and has a variable lenght because it's not necessary inform it's size at the beginning and here we are creating a empty list of maps.
//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets %T\n", conferenceName, conferenceTickets, remainingTickets)  // The %T can be used to see the value of a variable.

type UserData struct { // Here we are creating a struct, listing all the properties that these type should have, it's like a java class.
	firstName   string
	lastName    string
	email       string
	userTickets int
}

var wg = sync.WaitGroup{} // It waits for the launched goroutine to finish.

func main() { // The main function is where the code starts to run, without it, the programme won't know were to start.
	// go run 'main.go' is the command to run a programme in go.

	greetUsers() // Here, I am calling the defined function and passing the conference name as a parameter.

	//for { // It's used as a infinit loop.

	firstName, lastName, email, userTickets := getUserInput()

	//isValidCity := city == "Singapore" || city == "London"
	// !isValidCity  By doing that, we negate the variable.

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber { // It's a simple if/else statment.

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)                                              // It sets the number of goroutines to wait for.
		go sendTicket(userTickets, firstName, lastName, email) // By using the word 'go' in front of it, we are starting a new goroutine.

		firstNames := getFirstNames() // As the function returns just the first name, we can create a variable on the main function to get this return and use it anywhere.
		fmt.Printf("These first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 { // It's a simple if statment.
			fmt.Println("Our conference is booked out. Come back next year.")
			//break // It's used as a stop order.
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email adress you entered doesn't contain @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
		//continue  // It goes to next iteration
	}

	wg.Wait() // It blocks until the WaitGroup counter is 0.

	//}

	/*city := "London"

	switch city {
		case "New York":
			// execute code for booking New York conference tickets
		case "Singapore":
			// execute code for booking Singapore conference tickets
		case "London", "Netherlands":
			// execute code for booking London and Netherlands conference tickets, because they have the same type of conference tickets
		case "Berlin":
			// execute code for booking Berlin conference tickets
		case "Mexico city":
			// execute code for booking Mexico city conference tickets
		case "Hong Kong":
			// execute code for booking Hong Kong conference tickets
		default:
			fmt.Print("No valid city selected")
	}*/

}

func greetUsers() { // I created these function and I am passing an parameter and it's type.
	fmt.Printf("Welcome to %v booking application!\n", conferenceName) // It's the formated print, makes easier to writte things as below.
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string { // If I want to return something to the main function, I have to define the type of the return.
	firstNames := []string{}
	for _, booking := range bookings { // This '_' is used to ignore a varible that I don't want to use.
		firstNames = append(firstNames, booking.firstName)
		/*var names = strings.Fields(booking) // This function splits the string with white spaces as separator and returns a slice with the split elements ["Guilherme","Duran"].
		firstNames = append(firstNames, names[0]) */
	}
	return firstNames // The function returns the first name.
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName  This is how we would do if we had to work with an array, and as we don´t know the exactly number of people that are interested in the tickets and the first empty position to add this person, it's easier to use a slice.

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	/*var userData = make(map[string]string) // Here we are declaring an empty map, and we need to pass the type of the keys and the type of the values and both of them must have the same type. [string]int is a possible option.
	userData["firstName"] = firstName      // This is the first key/value map that we are saving in the variable.
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatInt(int64(userTickets), 10) // Here we are converting a integer to a string in base 10. */

	bookings = append(bookings, userData) // This is how we add someone to the slice.
	fmt.Printf("List of bookings is %v\n", bookings)
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
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)                                                       // This function give a delay to the application for 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // We are saving a formated sentence into a variable.
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done() // It says when the goroutine is finished, decrementing tha WaitGroup by 1.
}
