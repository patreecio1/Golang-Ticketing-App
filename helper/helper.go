package helper 

import 
"strings"


func Inputvalidation(firstname string, lastname string, email string,  usertickets uint, remainingtickets uint)(bool, bool, bool){
	isVlaidName := len(firstname) >= 2 && len(lastname) >= 2
   isValidEmail := strings.Contains(email, "@")
   isValidTicketNumber := usertickets > 0 && usertickets <= remainingtickets
   return isVlaidName, isValidEmail, isValidTicketNumber
}
