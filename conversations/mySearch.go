package conversations

import (
	"fmt"
	"log"
	"meet_bot/dbfellow"
	"meet_bot/keyboards"
)

//gets UserData struct, initiates search in database, results rating and sorting, deletes message,
//if nothing is found, sends an appropriate message to user.
//if something found, inserts a list of found IDs ordered from the most high rated to the least into Users (SearchResults),
//gets the first profile from Users, parse it into a HTML template
//and sends it to the user.
func Search(chatID int64, tgID string) error {

	userData := dbfellow.GetUserData(tgID)
	if userData.TgID == "" {
		err := fmt.Errorf("No userdata for search, tgID %s", tgID)
		log.Print(err)
		return err
	}

	people := dbfellow.SelectFromDBbySexAndAge(tgID, userData)
	alreadyInteractedWith := dbfellow.FindLikesFrom(tgID)
	newPeople := dbfellow.ExcludeLikes(alreadyInteractedWith, people)
	rated := dbfellow.RatingCalculation(userData, newPeople)
	list := dbfellow.SearchResultSorter(rated)

	if len(list) != 0 {

		err := dbfellow.UpdateTableWithStringArray("users", "searchResults", tgID, list)
		if err != nil {
			return err
		}

		ShowCV(list[0], chatID)

	} else {

		MessageCompiler(chatID, NothingFound, keyboards.KNothingFound)
	}
	return nil
}

func ShowCV(userID string, chatID int64) {

	ud := dbfellow.GetUserData(userID)
	text, photo := UserProfileTemplate(ud, false, chatID)
	keyboard := AddToOneButtonKeyboard(userID, keyboards.CV)
	MessageWithPhoto(text, photo, keyboard)
}
