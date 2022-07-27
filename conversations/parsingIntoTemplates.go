package conversations

import (
	"fmt"
	"log"
	"meet_bot/dbfellow"

	"meet_bot/templates"

	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//takes UserData struct and parse its' fields into UserProfile HTML template. Decodes photo into FileBite.
func UserProfileTemplate(ud dbfellow.UserData, edit bool, chatID int64) (string, tgbotapi.PhotoConfig) {

	var text string
	if edit {
		text = fmt.Sprintf(templates.EditUserProfile, ud.NameInBot, ud.Sex, ud.Age, ud.Location, ud.AboutMe, ud.Interests, ud.PartnersSex, ud.PartnersSex)
	} else {
		text = fmt.Sprintf(templates.UserProfile, ud.NameInBot, ud.Sex, ud.Age, ud.Location, ud.AboutMe, ud.Interests)
	}

	photo := ConvertStringToPhoto(ud.Photo, chatID)
	photoConf := tgbotapi.NewPhoto(chatID, photo)

	return text, photoConf

}

func CapitalizingDestinations(dests string) string {

	trimmed := strings.ToLower(dests)

	replace := func(word string) string {

		switch word {
		case "и", "на", "округ", "городской", "острова", "д":
			return word
		case "оаэ", "кндр", "сша":
			return strings.ToUpper(word)
		}
		return strings.Title(word)
	}

	r, err := regexp.Compile(`[а-яА-Я]+`)
	if err != nil {
		log.Printf("Error compiling regexp in func CapitalizingDestinations: %v", err)
	}
	dests2 := r.ReplaceAllStringFunc(trimmed, replace)
	return dests2
}
