package conversations

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"meet_bot/keyboards"

	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PassFeedback(message string, chatID int64) error {

	keyboard := AddToOneButtonKeyboard(chatID, keyboards.AnswerFeedback)
	str := fmt.Sprintf("<b>💌 Новый фидбек: </b>\n%s", message)
	receiver, _ := strconv.Atoi(os.Getenv("FeedbackReceiver"))
	msg := tgbotapi.NewMessage(int64(receiver), str)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboard

	_, err := Bot.Send(msg)
	if err != nil {
		err := fmt.Errorf("Can't send feedback '%s' from chat %v: %w", message, chatID, err)
		log.Print(err)
		return err
	}
	return nil
}

func AddToOneButtonKeyboard(ID interface{}, keyboard [][]string) tgbotapi.InlineKeyboardMarkup {

	str := fmt.Sprintf("%v||", ID)

	if len(strings.Split(keyboard[0][1], "||")) == 1 {

		keyboard[0][1] = str + keyboard[0][1]

	}

	newKeyboard := keyboards.KeyboardConstructor(1, keyboard)
	return newKeyboard
}
