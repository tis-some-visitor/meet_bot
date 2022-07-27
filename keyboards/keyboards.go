package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func KeyboardConstructor(columns int, content [][]string) tgbotapi.InlineKeyboardMarkup {

	var keyboard [][]tgbotapi.InlineKeyboardButton
	var row []tgbotapi.InlineKeyboardButton

	colCounter := columns
	buttons := ButtonsMaker(content)

	for _, button := range buttons {

		row = append(row, button)
		colCounter--
		if colCounter == 0 {
			keyboard = append(keyboard, row)
			row = nil
			colCounter = columns
		}
	}

	if len(row) > 0 {
		keyboard = append(keyboard, row)
	}

	var markup tgbotapi.InlineKeyboardMarkup

	markup.InlineKeyboard = keyboard

	return markup
}

func ButtonsMaker(texts [][]string) []tgbotapi.InlineKeyboardButton {

	var buttons []tgbotapi.InlineKeyboardButton

	for _, button := range texts {
		text := button[0]
		callback := button[1]
		b := tgbotapi.NewInlineKeyboardButtonData(text, callback)
		buttons = append(buttons, b)
	}

	return buttons
}

var (

	// edit destination

	EditWhere = [][]string{
		{"⬅️ Назад", "editDestinationBack"},
	}

	KEditWhere = KeyboardConstructor(1, EditWhere)

	//partner's sex

	PartnersSex = [][]string{
		{"Ж", "psЖ"},
		{"М", "psМ"},
		{"⬅️ Назад", "psBack"},
		{"Не важно 🗿", "psне важно"},
	}

	KPartnersSex = KeyboardConstructor(2, PartnersSex)

	EditPartnersSex = [][]string{
		{"Ж", "epsЖ"},
		{"М", "epsМ"},
		{"⬅️ Назад", "epsBack"},
		{"Не важно 🗿", "epsне важно"},
	}

	KEditPartnersSex = KeyboardConstructor(2, EditPartnersSex)

	PartnersAge = [][]string{
		{"🔲 18-24", "18-24"},
		{"🔲 25-34", "25-34"},
		{"🔲 35-44", "35-44"},
		{"🔲 45-54", "45-54"},
		{"🔲 55+", "55+"},
		{"Не важно!🤘🏽", "pAgeне важно"},
		{"⬅️ Назад", "partnersAgeBack"},
		{"Готово⛱", "savePartnersAge"},
	}

	KPartnersAge = KeyboardConstructor(2, PartnersAge)

	// edit short trip: partner's age

	EditPartnersAge = [][]string{
		{"🔲 18-24", "18-24"},
		{"🔲 25-34", "25-34"},
		{"🔲 35-44", "35-44"},
		{"🔲 45-54", "45-54"},
		{"🔲 55+", "55+"},
		{"Не важно!🤘🏽", "epAgeне важно"},
		{"⬅️ Назад", "ePartnersAgeBack"},
		{"Готово⛱", "esavePartnersAge"},
	}

	KEditPartnersAge = KeyboardConstructor(2, EditPartnersAge)

	//user's profile

	UserName = [][]string{
		{"Использовать имя из профиля телеграм", "useTGname"},
	}

	EditUserName = [][]string{
		{"Использовать имя из профиля телеграм", "euseTGname"},
	}

	UserSex = [][]string{
		{"Ж", "uЖ"},
		{"М", "uМ"},
		{"⬅️ Назад", "userSexBack"},
		{"другой/не указывать", "uне указан"},
	}

	EditUserSex = [][]string{
		{"Ж", "euЖ"},
		{"М", "euМ"},
		{"⬅️ Назад", "euserSexBack"},
		{"другой/не указывать", "euне указан"},
	}

	UserAge = [][]string{
		{"18-24", "U18-24"},
		{"25-34", "U25-34"},
		{"35-44", "U35-44"},
		{"45-54", "U45-54"},
		{"55+", "U55+"},
		{"⬅️ Назад", "userAgeBack"},
	}

	EditUserAge = [][]string{
		{"18-24", "eU18-24"},
		{"25-34", "eU25-34"},
		{"35-44", "eU35-44"},
		{"45-54", "eU45-54"},
		{"55+", "eU55+"},
		{"⬅️ Назад", "euserAgeBack"},
	}

	Interests = [][]string{
		{"🔲 культура", "культура"},
		{"🔲 психология", "психология"},

		{"🔲 экстрим", "экстрим"},
		{"🔲 природа", "природа"},

		{"🔲 ЗОЖ", "ЗОЖ"},
		{"🔲 анти-ЗОЖ", "анти-ЗОЖ"},

		{"🔲 еда", "еда"},
		{"🔲 спорт", "спорт"},

		{"🔲 путешествия", "путешествия"},
		{"⬅️ Назад", "interestsBack"},
		{"🔲 гейминг", "гейминг"},

		{"Готово!⛱", "saveInterests"},
	}

	EditInterests = [][]string{
		{"🔲 культура", "культура"},
		{"🔲 психология", "психология"},

		{"🔲 экстрим", "экстрим"},
		{"🔲 природа", "природа"},

		{"🔲 ЗОЖ", "ЗОЖ"},
		{"🔲 анти-ЗОЖ", "анти-ЗОЖ"},

		{"🔲 еда", "еда"},
		{"🔲 спорт", "спорт"},

		{"🔲 путешествия", "путешествия"},

		{"⬅️ Назад", "einterestsBack"},
		{"🔲 гейминг", "гейминг"},
		{"Готово!⛱", "esaveInterests"},
	}

	AboutMe = [][]string{
		{"⬅️ Назад", "aboutMeBack"},
		{"Пропустить ➡️", "skipAboutMe"},
	}

	EditAboutMe = [][]string{
		{"⬅️ Назад", "eaboutMeBack"},
	}

	Photo = [][]string{
		{"⬅️ Назад", "photoBack"},
	}

	EditPhoto = [][]string{
		{"⬅️ Назад", "ephotoBack"},
	}

	KUserName  = KeyboardConstructor(1, UserName)
	KUserSex   = KeyboardConstructor(2, UserSex)
	KUserAge   = KeyboardConstructor(3, UserAge)
	KInterests = KeyboardConstructor(3, Interests)
	KAboutMe   = KeyboardConstructor(2, AboutMe)
	KPhoto     = KeyboardConstructor(1, Photo)

	KEditUsername  = KeyboardConstructor(1, EditUserName)
	KEditUserSex   = KeyboardConstructor(1, EditUserSex)
	KEditUserAge   = KeyboardConstructor(3, EditUserAge)
	KEditInterests = KeyboardConstructor(3, EditInterests)
	KEditAboutMe   = KeyboardConstructor(1, EditAboutMe)
	KEditPhoto     = KeyboardConstructor(1, EditPhoto)

	//preview

	Preview = [][]string{
		{"Посмотреть", "preview"},
	}

	KPreview = KeyboardConstructor(1, Preview)

	EditOrSearch = [][]string{
		{"Редактировать", "editUserinfoOnPreview"},
		{"Искать 🧨", "go search"},
	}

	KEditOrSearch = KeyboardConstructor(1, EditOrSearch)

	Search = [][]string{
		{"ОК", "search"},
	}

	KSearch = KeyboardConstructor(1, Search)

	SearchRetry = [][]string{
		{"Искать!🧨", "search"},
	}

	KSearchRetry = KeyboardConstructor(1, SearchRetry)

	EditUserinfo = [][]string{
		{"Фото: загрузить другое", "editPhoto"},
		{"Имя: редактировать", "editName"},
		{"Пол: редактировать", "editSex"},
		{"Возраст: редактировать", "editAge"},
		{"Обо мне: редактировать", "editAboutMe"},
		{"Интересы: редактировать", "editInterests"},
		{"Пол попутчика: редактировать", "editPsex"},
		{"Возраст попутчика: редактировать", "editPage"},
		{"Сохранить изменения", "previewAfterEditUser"},
	}

	KEditUserinfo = KeyboardConstructor(1, EditUserinfo)

	NothingFound = [][]string{
		{"Изменить параметры поиска", "preview"},
	}

	KNothingFound = KeyboardConstructor(1, NothingFound)

	//VIEWING SEARCH RESULTS

	CV = [][]string{
		{"❤️", "like"},
		{"Пропустить ➡️", "nextCV"},
	}

	KCV = KeyboardConstructor(1, CV)

	NextCV = [][]string{
		{"Дальше ➡️", "nextCV"},
	}

	KNextCV = KeyboardConstructor(1, NextCV)

	EndOfFoundList = [][]string{
		{"Посмотреть заново", "watchCVsAgain"},
	}

	KEndOfFoundList = KeyboardConstructor(1, EndOfFoundList)

	//LIKE & MATCH

	SendMyUsername = [][]string{
		{"Отправить мой юзернейм", "sendMyUsername"},
	}

	KSendMyUsername = KeyboardConstructor(1, SendMyUsername)

	IWasLiked = [][]string{
		{"Отправить мой юзернейм", "sendMyUsername"},
		{"Дальше ➡️", "nextCV"},
	}

	KIWasLiked = KeyboardConstructor(1, IWasLiked)

	EnteredUsername = [][]string{
		{"Заполнил юзернейм, повтори отправку", "repeatSendingUsername"},
		{"Скрыть", "delete"},
	}

	KEnteredUsername = KeyboardConstructor(1, EnteredUsername)

	//MY PROFILE

	MyProfile = [][]string{
		{"Редактировать", "editUserinfoGeneral"},
		{"Скрыть", "delete"},
	}

	KMyProfile = KeyboardConstructor(1, MyProfile)

	//MY REQUESTS

	LikeView = [][]string{
		{"Отправить мой юзернейм", "sendMyUsername"},
		{"Следующий ➡️", "showNextRequest"},
	}

	KLikeView = KeyboardConstructor(1, LikeView)

	//MY MATCHES

	MatchView = [][]string{
		{"Следующий ➡️", "showNextMatch"},
	}

	KMatchView = KeyboardConstructor(1, MatchView)

	//other

	EmptyKeyboard = tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
	}

	Restart = [][]string{
		{"СТАРТ", "newTrip"},
	}

	KRestart = KeyboardConstructor(1, Restart)

	Hide = [][]string{
		{"Скрыть", "delete"},
	}

	KHide = KeyboardConstructor(1, Hide)

	AnswerFeedback = [][]string{
		{"Ответить", "answerFeedback"},
	}

	KAnswerFeedback = KeyboardConstructor(1, AnswerFeedback)

	Delete = [][]string{
		{"OK", "delete"},
	}

	KDelete = KeyboardConstructor(1, Delete)
)
