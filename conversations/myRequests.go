package conversations

import (
	"meet_bot/dbfellow"
	"meet_bot/keyboards"
)

//saves given slice of IDs to Users, calls for ShowNextRequest
func ShowMyRequests(IDs []string, chatID int64, tgID string) error {

	err := dbfellow.UpdateTableWithStringArray("users", "likeIDs", tgID, IDs)
	if err != nil {
		return err
	}

	ShowRequest(chatID, IDs[0])
	return err
}

func ShowRequest(chatID int64, ID string) error {

	ud := dbfellow.GetUserData(ID)
	userCV, userPhoto := UserProfileTemplate(ud, false, chatID)
	text := MyRequestsHeader + userCV
	keyboardWithLikeID := AddToOneButtonKeyboard(ID, keyboards.LikeView)

	_, err := MessageWithPhoto(text, userPhoto, keyboardWithLikeID)
	return err
}
