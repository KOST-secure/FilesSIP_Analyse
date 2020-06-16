package main

import (
	"fmt"
)

// Simple individual tasks
func makeHotelReservation() {
	fmt.Println("Done making hotel reservation.")
}
func bookFlightTickets() {
	fmt.Println("Done booking flight tickets.")
}
func orderADress() {
	fmt.Println("Done ordering a dress.")
}
func payCreditCardBills() {
	fmt.Println("Done paying Credit Card bills.")
}

// Tasks that will be executed in parts

// Writing Mail
func writeAMail() {
	fmt.Println("Wrote 1/3rd of the mail.")
	continueWritingMail1()
}
func continueWritingMail1() {
	fmt.Println("Wrote 2/3rds of the mail.")
	continueWritingMail2()
}
func continueWritingMail2() {
	fmt.Println("Done writing the mail.")
}

// Listening to Audio Book
func listenToAudioBook() {
	fmt.Println("Listened to 10 minutes of audio book.")
	continueListeningToAudioBook()
}
func continueListeningToAudioBook() {
	fmt.Println("Done listening to audio book.")
}

// All the tasks we want to complete in the day.
// Note that we do not include the sub tasks here.
var listOfTasks = []func(){
	makeHotelReservation, bookFlightTickets, orderADress,
	payCreditCardBills, writeAMail, listenToAudioBook,
}

func main() {
	for _, task := range listOfTasks {
		task()
	}
}
