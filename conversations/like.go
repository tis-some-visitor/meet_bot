package conversations

import (
	"meet_bot/dbfellow"
	"meet_bot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//updates saved username, sends notion about new like with CV to liked user,
//deletes liked CV message, send "like was/wasn't send" to user
func Like(update tgbotapi.Update, userTgID string, likeID int, myUsername string, myTgID string) {

	dbfellow.UpdateTableWithText("states_", "title", myUsername, myTgID)

	adressChat, _ := dbfellow.GetUsernameAndChatID(userTgID)
	myUserData := dbfellow.GetUserData(myTgID)

	myCV, myPhoto := UserProfileTemplate(myUserData, false, adressChat)
	likeMessageText := LikeMessageHeader + myCV

	DeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	keyboard := AddToOneButtonKeyboard(likeID, keyboards.SendMyUsername)

	_, err := MessageWithPhoto(likeMessageText, myPhoto, keyboard)
	if err != nil {
		MessageCompiler(update.CallbackQuery.Message.Chat.ID, QLikeWasntSend, keyboards.KNextCV)
	} else {
		ExcludeFromSearchResults(userTgID, myTgID)
		MessageCompiler(update.CallbackQuery.Message.Chat.ID, QLikeSend, keyboards.KNextCV)
	}
}

func ExcludeFromSearchResults(ID string, tgID string) {

	dbfellow.DeleteStringFromArray("users", "searchResults", ID, tgID)
	dbfellow.SubstructOneFromIndexColumn("users", "currentCV", tgID)
}
