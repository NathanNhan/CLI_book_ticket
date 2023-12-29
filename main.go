package main

import (
	"fmt"
	"strings"
)

func main() {
	//Khai báo biến
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	//Array in golang
	// bookings := [50]string{}
	//Slice in golang , có thể mở rộng dc kích thước mảng
	bookings := []string{}
	//Giá trị của biến không được là số âm
	var remainingTickets uint = 50
	//Greeting
	greeting(conferenceName, conferenceTickets, remainingTickets)
	//Nhập xuất dữ liệu bằng Scan
	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("Enter your Fist name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email: ")
		fmt.Scan(&email)

		fmt.Println("Enter your User's ticket: ")
		fmt.Scan(&userTicket)
		//Validate input data
		validateName, validateEmail, valiedateUserTicket := validate(firstName, lastName, email, userTicket, remainingTickets)
		//Add value to Array := bookings
		// bookings[0] = firstName + " " + lastName
		//Add values to slice
		if validateName && validateEmail && valiedateUserTicket {
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+lastName)
			// fmt.Printf("Array %v \n", bookings)
			// fmt.Printf("Bookings Array of First Element %v \n", bookings[0])

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("Remaining booking are %d\n", remainingTickets)

			//Lấy firstname trong mảng bookings
			firstNames := printFirstName(bookings)
			fmt.Printf("Booking list %v \n", firstNames)

			//Nếu remainbooking = 0 => break
			if remainingTickets == 0 {
				fmt.Println("All Tickets booked out. Coming next year")
				break
			}
		} else {
			if !validateName {
				fmt.Println("Your FirstName or Last name at least 2 characters")
			}
			if !validateEmail {
				fmt.Println("Your email must be contain @")
			}
			if !valiedateUserTicket {
				fmt.Println("Your ticket must greater > 0")
			}
		}

	}
}

// Hàm greeting
func greeting(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickers remaining are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

// Hàm print first name
func printFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// Hàm validate
func validate(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	validateName := len(firstName) >= 2 && len(lastName) >= 2
	validateEmail := strings.Contains(email, "@")
	valiedateUserTicket := userTicket > 0 && userTicket <= remainingTickets
	return validateName, validateEmail, valiedateUserTicket
}
