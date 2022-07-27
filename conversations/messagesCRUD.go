package conversations

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//creates a message with parse mode set to HTML and with inline keyboard attached, sends it to the given chat.
func MessageCompiler(chatID int64, text string, keyboard tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"

	message, err := Bot.Send(msg)
	if err != nil {
		pieceOfText := text[0 : len(text)/2]
		err := fmt.Errorf("Can't send message \"%s\"... to chat %v: %w", pieceOfText, chatID, err)
		log.Print(err)
		return message, err
	}

	return message, err
}

//creates a new message as photo with given text (parse mode set to HTML) and inline keyboard, sends it to user
func MessageWithPhoto(text string, photo tgbotapi.PhotoConfig, keyboard tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {

	photo.Caption = text
	photo.ParseMode = "HTML"
	photo.ReplyMarkup = keyboard

	message, err := Bot.Send(photo)
	if err != nil {
		pieceOfText := text[0 : len(text)/2]
		err := fmt.Errorf("Can't send photo with message \"%s\"... to chat %v: %w", pieceOfText, message.Chat.ID, err)
		log.Print(err)
	}

	return message, err
}

//edits a callback message and an inline keyboard attached
func MessageEditor(u tgbotapi.Update, text string, keyboard tgbotapi.InlineKeyboardMarkup) {

	msg := tgbotapi.NewEditMessageTextAndMarkup(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID, text, keyboard)
	msg.ParseMode = "HTML"

	if _, err := Bot.Send(msg); err != nil {
		log.Printf("Can't edit message %v in chat %v: %v", u.CallbackQuery.Message.MessageID, u.CallbackQuery.Message.Chat.ID, err)
	}
}

//deletes message given in callback and sends the next one
func MessageUpdater(u tgbotapi.Update, text string, keyboard tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {

	DeleteMessage(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID)

	msg := tgbotapi.NewMessage(u.CallbackQuery.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"

	message, err := Bot.Send(msg)
	if err != nil {
		pieceOfText := text[0 : len(text)/2]
		err := fmt.Errorf("Can't send message \"%s\"... to chat %v: %w", pieceOfText, u.CallbackQuery.Message.Chat.ID, err)
		log.Print(err)
	}

	return message, err
}

func MessageCompilerFeedbackReply(replyTo string, chatID int64) error {

	txt := fmt.Sprintf("Напишите Ваш ответ на фидбек \"%s\": ", replyTo)
	msg := tgbotapi.NewMessage(chatID, txt)

	_, err := Bot.Send(msg)
	if err != nil {
		err := fmt.Errorf("Can't send reply to feedback \"%s\" to chat %v: %w", replyTo, chatID, err)
		log.Print(err)
		return err
	}
	return err
}

//edits a given message and its inline keyboard
func MessageEditorIDSpecific(chatID int64, messageID int, text string, keyboard tgbotapi.InlineKeyboardMarkup) {

	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, messageID, text, keyboard)
	msg.ParseMode = "HTML"
	if _, err := Bot.Send(msg); err != nil {
		panic(err)
	}
}

//replaces a given message and its inline keyboard
func MessageUpdaterIDSpecific(u tgbotapi.Update, messageID int, text string, keyboard tgbotapi.InlineKeyboardMarkup) {

	DeleteMessage(u.Message.Chat.ID, messageID)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"

	_, err := Bot.Send(msg)
	if err != nil {
		pieceOfText := text[0 : len(text)/2]
		err := fmt.Errorf("Can't send message \"%s\"... to chat %v: %w", pieceOfText, u.Message.Chat.ID, err)
		log.Print(err)
	}
}

func DeleteMessage(chatID int64, messageID int) {

	deleteMessage := tgbotapi.NewDeleteMessage(chatID, messageID)

	_, err := Bot.Request(deleteMessage)
	if err != nil {
		log.Printf("Can't delete message %d in chat %d: %s", messageID, chatID, err)
	}

	log.Printf("Deleted message %d in chat %d", messageID, chatID)
}

//sends error message to user, waits 1 second, deletes user's message, deletes error message
func ErrorAndWipeBothMessages(u tgbotapi.Update, errorMessage string) {

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, errorMessage)
	foo, err := Bot.Send(msg)
	if err != nil {
		log.Print(err)
	}

	time.Sleep(1 * time.Second)

	deleteUserMessage := tgbotapi.NewDeleteMessage(u.Message.Chat.ID, u.Message.MessageID)
	_, err = Bot.Request(deleteUserMessage)
	if err != nil {
		log.Print(err)
	}

	deleteOurMessage := tgbotapi.NewDeleteMessage(u.Message.Chat.ID, foo.MessageID)
	_, err = Bot.Request(deleteOurMessage)
	if err != nil {
		log.Print(err)
	}
}

//sends info message to user, waits 1 second, deletes info message
func PopUpInfoMessage(chatID int64, text string) {

	msg := tgbotapi.NewMessage(chatID, text)
	foo, err := Bot.Send(msg)
	if err != nil {
		log.Printf("Can't send pop-up %s to chat %v: %v", text, chatID, err)
	}

	time.Sleep(1 * time.Second)

	deleteOurMessage := tgbotapi.NewDeleteMessage(chatID, foo.MessageID)
	_, err = Bot.Request(deleteOurMessage)
	if err != nil {
		log.Printf("Can't delete pop-up %s in chat %v: %v", text, chatID, err)
	}
}

//sends given message if err != nil
func SendIfProblem(chatID int64, text string, keyboard tgbotapi.InlineKeyboardMarkup, err error) {

	if err != nil {

		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = keyboard

		_, err := Bot.Send(msg)
		if err != nil {
			log.Printf("Can't send a problem message %s to chat %v: %v", text, chatID, err)
		}
	}
}
