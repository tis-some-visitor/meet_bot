package conversations

import (
	"meet_bot/dbfellow"
	"meet_bot/maps"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DestinationValidAndSave(update tgbotapi.Update, edit bool, tgID string, keyboardMesNO int,
	messageValid string, keyboardValid tgbotapi.InlineKeyboardMarkup) {

	valid, validPlace := GeographycallyValid(update.Message.Text)

	if valid {

		SaveDestination(validPlace, update, tgID)
		dbfellow.SetUserStateToNeutral(tgID)
		if edit {
			BackToEditedUserinfoMessage(update, tgID, keyboardMesNO)
		} else {
			MessageUpdaterIDSpecific(update, keyboardMesNO, messageValid, keyboardValid)
		}

	} else {

		ErrorAndWipeBothMessages(update, "Локация не распознана, попробуйте еще раз")
	}
}

func GeographycallyValid(text string) (valid bool, place string) {

	dataPieceReady := strings.ToLower(strings.TrimSpace(text))
	_, okT := maps.RussianTowns[dataPieceReady]
	_, okC := maps.Countries[dataPieceReady]
	if okT || okC {
		valid = true
		place = CapitalizingDestinations(dataPieceReady)
	}
	return
}

func SaveDestination(validPlace string, update tgbotapi.Update, tgID string) {

	dbfellow.UpdateTableWithText("trips", "location", validPlace, tgID)
	DeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
}
