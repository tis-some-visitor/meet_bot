package conversations

import (
	"fmt"
	"meet_bot/dbfellow"
	"meet_bot/keyboards"
	"meet_bot/templates"
)

func ShowMyMatches(IDs []string, chatID int64, tgID string) error {

	err := dbfellow.UpdateTableWithStringArray("users", "matches", tgID, IDs)
	if err != nil {
		return err
	}

	ShowMatch(chatID, IDs[0], tgID)
	return err
}

func ShowMatch(chatID int64, ID string, tgID string) error {

	ud := dbfellow.GetUserData(ID)
	_, username := dbfellow.GetUsernameAndChatID(tgID)

	text, photo := UserProfileTemplate(ud, false, chatID)
	text = fmt.Sprintf(templates.ViewMatch, username) + text
	_, err := MessageWithPhoto(text, photo, keyboards.KMatchView)
	return err

}
