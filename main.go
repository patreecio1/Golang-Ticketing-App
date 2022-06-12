package main

import (
	"fmt"
	"TICKETING-APP/helper"
)
var cinemaname = "Lucky cinema"
const moviestickets int = 50
var remainingtickets uint = 50
 var booking = make([]userData, 0)

 type userData struct{
	 firstName string
	 lastName string
	 email string
	 userTickets uint
 }
 
func main(){
	
	 greetusers()
	for{

		firstname, lastname,email, usertickets:=userinput()

		isVlaidName, isValidEmail, isValidTicketNumber:= helper.Inputvalidation(firstname, lastname, email, usertickets, remainingtickets)
		if isVlaidName && isValidEmail && isValidTicketNumber {

			bookTickets(firstname, lastname,  usertickets, email)
		
	
		    FirstName:= getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", FirstName)
			if remainingtickets == 0{
				fmt.Printf("Movie Tickets sold out, please come back next year\n")
				break
			
		}
		}else {
			if !isVlaidName{
				fmt.Printf("Firstname or Lastname are too short\n")
			}
			if !isValidEmail{

































































































































































































































































				
				fmt.Printf("Your email is incorrect\n")
			}
			if !isValidTicketNumber{
				fmt.Printf("the Movie ticket number requested is already greater than the remaining tickets\n")
			} 
			
			
		}
		
	}
	
}
func greetusers(){

	fmt.Printf("Welcome to %v booking application for Amazing movies\n", cinemaname)
	fmt.Printf("We have total of %v tickets and %v are remaining\n",moviestickets, remainingtickets)
	fmt.Println("Get your movie tickets here to attend")

}

func getFirstNames() []string{
	FirstName := []string{}
	for _, booking:= range booking{
		FirstName = append(FirstName, booking.firstName)
	}
	return FirstName

} 

func userinput()(string, string, string, uint){
	var firstname string
	var  lastname string
	var usertickets uint
	var email string
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstname)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastname)
	fmt.Println("Please enter your email")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets needed")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets
}

func bookTickets(firstname string , lastname string , usertickets uint , email string, ){
	remainingtickets = remainingtickets - usertickets

	var userData = userData{
		firstName: firstname,
		lastName: lastname,
		email: email,
		userTickets: usertickets,

	}
	booking = append(booking, userData )
	fmt.Printf("THE NAME IS %v\n",booking)
	fmt.Printf("User %v %v booked %v movie tickets. You will receive a confirmation with this mail %v\n", firstname, lastname, usertickets,email)
	fmt.Printf("Remaining tickets %v for %v\n",remainingtickets, cinemaname)
	fmt.Printf("These are all our bookings %v\n",booking)

}