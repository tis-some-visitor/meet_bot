package conversations

import (
	"fmt"
	"meet_bot/dbfellow"
	"meet_bot/keyboards"
	"meet_bot/templates"
)

//if username of the user is set (not ""), sends his trip&userinfo with his username to other user.
//if username is "" sends message to user with a reminder to fill it in.
func SendUsername(myChatID int64, userID, myUsername, myTgID string) error {

	if myUsername != "" {

		userChatID, userUsername := dbfellow.GetUsernameAndChatID(userID)
		myUD := dbfellow.GetUserData(myTgID)
		myCV, myPhoto := UserProfileTemplate(myUD, false, userChatID)
		header := fmt.Sprintf(templates.NewMatch, myUsername)
		message := header + myCV

		_, err := MessageWithPhoto(message, myPhoto, keyboards.EmptyKeyboard)
		if err != nil {
			MessageCompiler(myChatID, CantSendUsername, keyboards.EmptyKeyboard)
		} else {
			_, _, likeID := dbfellow.AlreadyInLikes(userID, myTgID)
			dbfellow.Match(likeID, myTgID, myUsername)
			header := fmt.Sprintf(templates.ConfirmMatch, userUsername)
			hisUD := dbfellow.GetUserData(userID)
			hisCV, hisPhoto := UserProfileTemplate(hisUD, false, myChatID)
			message = header + hisCV
			MessageWithPhoto(message, hisPhoto, keyboards.EmptyKeyboard)
		}

	} else {
		keyboard := AddToOneButtonKeyboard(userID, keyboards.EnteredUsername)
		MessageCompiler(myChatID, NoUserName, keyboard)
		return nil
	}
	return nil
}
