package conversations

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//process callbacks from checkbox-type keyboards. It takes a keyboard from a message, goes through all the buttons and compair
//each of them to a callback data. When the button is found, it's emoji prefix (green tick or empty checkbox) is changed for
// the other, and new keyboard is used to edit the original message
func CheckboxCallback(u tgbotapi.Update) error {

	keyboard := u.CallbackQuery.Message.ReplyMarkup.InlineKeyboard

	for j, row := range u.CallbackQuery.Message.ReplyMarkup.InlineKeyboard {
		for k, button := range row {
			if *button.CallbackData == u.CallbackQuery.Data {
				if strings.HasPrefix(keyboard[j][k].Text, tick) {
					keyboard[j][k].Text = strings.ReplaceAll(keyboard[j][k].Text, tick, checkbox)
				} else {
					keyboard[j][k].Text = strings.ReplaceAll(keyboard[j][k].Text, checkbox, tick)
				}
			}
		}
	}

	msg := tgbotapi.NewEditMessageReplyMarkup(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: keyboard})
	_, err := Bot.Send(msg)
	return err
}

//takes a keyboard from a callback message, concates texts of the buttons marked by green ticks into a single string.
func CollectMultiOptions(u tgbotapi.Update) string {

	keyboard := u.CallbackQuery.Message.ReplyMarkup.InlineKeyboard

	var stringOfChoices string

	for _, row := range keyboard {
		for _, button := range row {

			option := string(button.Text)
			if strings.HasPrefix(option, tick) {
				bareOption := strings.TrimPrefix(option, tick)
				stringOfChoices = stringOfChoices + bareOption + ", "
			}
		}
	}
	tr := strings.TrimSuffix(stringOfChoices, ", ")

	return tr
}

func SomethingChosen(stringOfChoices string) bool {

	if len(stringOfChoices) == 0 {
		return false
	}
	return true
}
