package conversations

import (
	"fmt"
	"meet_bot/dbfellow"
	"meet_bot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//handling of user input according to the current user state
func UserStateHandler(update tgbotapi.Update, state string, keyboardMesNO int, tgID string) {

	switch state {

	//location
	case "location":

		DestinationValidAndSave(update, false, tgID, keyboardMesNO, QPartnersSex,
			keyboards.KPartnersSex)

	case "editLocation":

		DestinationValidAndSave(update, true, tgID, keyboardMesNO, "", keyboards.EmptyKeyboard)

	case "username":

		if SybmolsNumber(update.Message.Text, 50) == true {

			dbfellow.UpdateTableWithText("users", "nameInBot", update.Message.Text, tgID)
			dbfellow.SetUserStateToNeutral(tgID)
			DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			MessageEditorIDSpecific(update.Message.Chat.ID, keyboardMesNO, QUserSex, keyboards.KUserSex)

		} else {

			ErrorAndWipeBothMessages(update, SomthWrong)
		}

	case "editName":

		if SybmolsNumber(update.Message.Text, 50) == true {

			dbfellow.UpdateTableWithText("users", "nameInBot", update.Message.Text, tgID)
			dbfellow.SetUserStateToNeutral(tgID)
			BackToEditedUserinfoMessage(update, tgID, keyboardMesNO)

		} else {

			ErrorAndWipeBothMessages(update, SomthWrong)
		}

	case "aboutMe":

		if SybmolsNumber(update.Message.Text, 500) == true {

			dbfellow.UpdateTableWithText("users", "aboutMe", update.Message.Text, tgID)
			dbfellow.SetUserState("photo", keyboardMesNO, tgID)
			DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			MessageEditorIDSpecific(update.Message.Chat.ID, keyboardMesNO, QPhoto, keyboards.KPhoto)

		} else {

			ErrorAndWipeBothMessages(update, SomthWrong)
		}

	case "editAboutMe":

		if SybmolsNumber(update.Message.Text, 500) == true {

			dbfellow.UpdateTableWithText("users", "aboutMe", update.Message.Text, tgID)
			BackToEditedUserinfoMessage(update, tgID, keyboardMesNO)

		} else {

			ErrorAndWipeBothMessages(update, SomthWrong)
		}

	case "photo":
		if update.Message.Document != nil || update.Message.Photo != nil {

			PhotoHandling(update, tgID, keyboardMesNO)

		} else {
			ErrorAndWipeBothMessages(update, SomthWrong)
		}

	case "editPhoto":

		if update.Message.Document != nil || update.Message.Photo != nil {

			PhotoHandlingEditing(update, tgID, keyboardMesNO)

		} else {

			ErrorAndWipeBothMessages(update, SomthWrong)
		}

	case "feedback":
		err := PassFeedback(update.Message.Text, update.Message.Chat.ID)
		dbfellow.SetUserStateToNeutral(tgID)
		if err != nil {
			MessageCompiler(update.Message.Chat.ID, "Отправить фидбек не удалось. Пожалуйста, попробуйте чуть позже", keyboards.EmptyKeyboard)
		} else {
			MessageCompiler(update.Message.Chat.ID, QLikeSend, keyboards.EmptyKeyboard)
		}

	case "feedbackReply":
		msg := fmt.Sprintf("💌 Получен новый ответ на фидбек: %s", update.Message.Text)
		_, err := MessageCompiler(int64(keyboardMesNO), msg, keyboards.EmptyKeyboard)
		dbfellow.SetUserStateToNeutral(tgID)
		if err != nil {
			MessageCompiler(update.Message.Chat.ID, "Отправка не удалась. Пожалуйста, попробуйте позже", keyboards.EmptyKeyboard)
		} else {
			MessageCompiler(update.Message.Chat.ID, QLikeSend, keyboards.EmptyKeyboard)
		}

	default:

		ErrorAndWipeBothMessages(update, UnknownCommand)
	}

}
