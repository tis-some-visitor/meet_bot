package conversations

import (
	"log"
	"strconv"
)

//checks if a string is not empty and contains given number of characters or less.
//return false if string length is 0 or more than given length
func SybmolsNumber(text string, length int) bool {

	if len(text) <= length && len(text) > 0 {
		log.Print("len = ", len(text))
		return true
	}

	return false
}

//finds the leftmost number in a given string, if found - converts it to int, returns true and the found number.
//returns false and zero if nothing found
func ItIsANumber(message string) (bool, int) {

	found := regNumVerifier.FindString(message)

	if len(found) != 0 {
		tripID, _ := strconv.Atoi(found)
		return true, tripID
	} else {
		return false, 0
	}
}
