package conversations

import (
	"fmt"
	"log"
	"meet_bot/dbfellow"
	"meet_bot/envar"
	"meet_bot/evil"
	"meet_bot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(chatID, ID int64, messageID int, name string, tgID string) {

	err := AddUserToDatabase(tgID, name, ID)
	if err != nil {
		MessageCompiler(chatID, "Что-то пошло не так, давай попробуем еще раз!", keyboards.KRestart)
	} else {
		dbfellow.SetUserState("location", messageID, tgID)
		MessageCompiler(chatID, QStart, keyboards.EmptyKeyboard)
	}
}

//creates user's profile and state entry in the database (if not exist).
func AddUserToDatabase(tgID, firstname string, ID int64) error {

	err := dbfellow.CreateNewUser(tgID, firstname)
	if err != nil {
		err = fmt.Errorf("Can't create user, tgID %s: %s", tgID, err)
		log.Print(err)
		return err
	}

	err = dbfellow.CreateNewState(tgID, ID)
	if err != nil {
		err = fmt.Errorf("Can't create state, tgID %s: %s", tgID, err)
		log.Print(err)
		return err
	}

	return err
}

func HideID(ID int64) string {

	tgID, err := evil.Retry(envar.HashID, ID)
	if err != nil {
		log.Panicf("Can't hash: %s", err)
	}
	return tgID

}

//get user data, parse it into template with HTML tags, delete user's message and the previous (keyboarded) one,
//send new profile to the user
func BackToEditedUserinfoMessage(update tgbotapi.Update, tgID string, keyboardMesNo int) {

	chatID := update.Message.Chat.ID

	data := dbfellow.GetUserData(tgID)
	text, photo := UserProfileTemplate(data, true, chatID)

	DeleteMessage(chatID, update.Message.MessageID)
	DeleteMessage(chatID, keyboardMesNo)
	MessageWithPhoto(text, photo, keyboards.KEditUserinfo)
}

//extracts userinfo from table Users into UserData struct, parses it into HTML template, decodes user photo from database,
//deletes a callback message and calls for MessageWithPhoto to send photo + templated userinfo + editUserinfo keyboard to user.
func BackToEditedUserinfoMessageFromCallback(update tgbotapi.Update, keyboard tgbotapi.InlineKeyboardMarkup, tgID string) {

	chatID := update.CallbackQuery.Message.Chat.ID

	data := dbfellow.GetUserData(tgID)

	text, photo := UserProfileTemplate(data, true, chatID)

	DeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	MessageWithPhoto(text, photo, keyboard)
}

func PhotoHandling(update tgbotapi.Update, tgID string, keyboardMesNO int) {

	photo := Download(update)
	dbfellow.UpdateTableWithText("users", "photo", photo, tgID)
	dbfellow.SetUserStateToNeutral(tgID)
	dbfellow.SavedSwitch(true, tgID)

	DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)

	MessageUpdaterIDSpecific(update, keyboardMesNO, QReady, keyboards.KPreview)

}

func PhotoHandlingEditing(update tgbotapi.Update, tgID string, keyboardMesNO int) {

	photo := Download(update)
	dbfellow.UpdateTableWithText("users", "photo", photo, tgID)
	dbfellow.SavedSwitch(true, tgID)
	BackToEditedUserinfoMessage(update, tgID, keyboardMesNO)
}

func ShowOrReportEmpty(whatToShow string, slice []string, count int, msgEmpty, msgError string, chatID int64, tgID string) {

	var show func([]string, int64, string) error

	switch whatToShow {
	case "requests":
		show = ShowMyRequests
	case "matches":
		show = ShowMyMatches
	}

	if count > 0 {
		err := show(slice, chatID, tgID)
		SendIfProblem(chatID, msgError, keyboards.EmptyKeyboard, err)
	} else {
		MessageCompiler(chatID, msgEmpty, keyboards.EmptyKeyboard)
	}
}
