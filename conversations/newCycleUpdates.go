package conversations

import (
	"log"
	"meet_bot/dbfellow"
	"meet_bot/keyboards"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func NewCycleUpdates(update tgbotapi.Update) {

	var tgID string

	if update.Message != nil {

		tgID = HideID(update.Message.From.ID)
		ID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		messageID := update.Message.MessageID
		firstName := update.Message.From.FirstName

		if update.Message.Text == "/start" {

			DeleteMessage(chatID, messageID)
			Start(chatID, ID, messageID, firstName, tgID)

		} else if update.Message.Text == "/search" {

			if dbfellow.UserProfileCompleted(tgID) {
				Search(chatID, tgID)
			} else {
				Start(chatID, ID, messageID, firstName, tgID)
			}

		} else if update.Message.Text == "/my_profile" {

			userInfo := dbfellow.GetUserData(tgID)
			text, photo := UserProfileTemplate(userInfo, true, chatID)
			MessageWithPhoto(text, photo, keyboards.KMyProfile)

		} else if update.Message.Text == "/requests" {

			IDs := dbfellow.GetMyRequests(tgID)
			ShowOrReportEmpty("requests", IDs, len(IDs), NoLikesYet, ErrorIndef, chatID, tgID)

		} else if update.Message.Text == "/matches" {

			IDs := dbfellow.GetMyMatches(tgID)
			ShowOrReportEmpty("matches", IDs, len(IDs), NoMatchesYet, ErrorIndef, chatID, tgID)

		} else if update.Message.Text == "/feedback" {

			message, _ := MessageCompiler(chatID, FeedbackPlease, keyboards.EmptyKeyboard)
			dbfellow.SetUserState("feedback", message.MessageID, tgID)

		} else {

			userProfile, _ := dbfellow.EntriesWithGivenTgID(tgID, "users")

			if userProfile > 0 {
				state, messageNO := dbfellow.GetUserState(tgID)
				newUserStateHandler(update, state, messageNO, tgID)

			} else {
				Start(chatID, ID, messageID, firstName, tgID)
			}
		}

	} else if update.CallbackQuery != nil {

		tgID = HideID(update.CallbackQuery.From.ID)

		chatID := update.CallbackQuery.Message.Chat.ID
		messageID := update.CallbackQuery.Message.MessageID
		data := update.CallbackQuery.Data

		switch data {

		//EDIT location

		case "editDestination":
			dbfellow.SetUserState("location", messageID, tgID)
			MessageUpdater(update, QStart, keyboards.KEditWhere)

		case "editDestinationBack":
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

			//partner's sex

		case "psЖ", "psМ", "psне важно":
			text := strings.TrimPrefix(data, "ps")
			dbfellow.UpdateTableWithText("users", "partnersSex", text, tgID)
			MessageEditor(update, QPartnersAge, keyboards.KPartnersAge)

		case "psBack":
			MessageEditor(update, QStart, keyboards.EmptyKeyboard)

			//edit partner's sex

		case "editPsex":
			MessageUpdater(update, QEditPartnerSex, keyboards.KEditPartnersSex)

		case "epsЖ", "epsМ", "epsне важно":
			text := strings.TrimPrefix(data, "eps")
			dbfellow.UpdateTableWithText("users", "partnersSex", text, tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		case "epsBack":
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

			//PARTNER's AGE

		case "18-24", "25-34", "35-44", "45-54", "55+":
			err := CheckboxCallback(update)
			if err != nil {
				log.Printf("Problem with checkboxes, user %s: %v", tgID, err)
			}

		case "pAgeне важно":
			text := strings.TrimPrefix(data, "pAge")
			dbfellow.UpdateTableWithText("users", "partnersAge", text, tgID)
			dbfellow.SetUserState("name", messageID, tgID)
			MessageEditor(update, QEnterUserName, keyboards.KUserName)

		case "partnersAgeBack":
			MessageEditor(update, QPartnersSex, keyboards.KPartnersSex)

		case "savePartnersAge":
			choice := CollectMultiOptions(update)
			if SomethingChosen(choice) == true {
				dbfellow.UpdateTableWithText("users", "partnersAge", choice, tgID)
				dbfellow.SetUserState("name", messageID, tgID)
				MessageEditor(update, QEnterUserName, keyboards.KUserName)
			} else {
				PopUpInfoMessage(chatID, NothingWasChosen)
			}

		//edit PARTNER's AGE

		case "editPage":

			MessageUpdater(update, QPartnersAge, keyboards.KEditPartnersAge)

		case "epAgeне важно":
			text := strings.TrimPrefix(data, "epAge")
			dbfellow.UpdateTableWithText("users", "partnersAge", text, tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		case "eshortPartnersAgeBack":
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		case "esavePartnersAge":
			choice := CollectMultiOptions(update)
			if SomethingChosen(choice) == true {
				dbfellow.UpdateTableWithText("users", "partnersAge", choice, tgID)
				BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)
			} else {
				PopUpInfoMessage(chatID, NothingWasChosen)
			}

		// NAME

		case "useTGname":
			dbfellow.SetUserStateToNeutral(tgID)
			name := dbfellow.GetSavedData("users", "firstname", tgID)
			dbfellow.UpdateTableWithText("users", "nameInBot", name, tgID)
			MessageEditor(update, QUserSex, keyboards.KUserSex)

		// edit NAME

		case "editName":
			DeleteMessage(chatID, messageID)
			message, _ := MessageCompiler(chatID, QEnterUserName, keyboards.KEditUsername)
			dbfellow.SetUserState("editName", message.MessageID, tgID)

		case "euseTGname":
			dbfellow.SetUserStateToNeutral(tgID)
			name := dbfellow.GetSavedData("users", "firstname", tgID)
			dbfellow.UpdateTableWithText("users", "nameInBot", name, tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		//SEX

		case "uМ", "uЖ", "uне указан":
			text := strings.TrimPrefix(data, "u")
			dbfellow.UpdateTableWithText("users", "sex", text, tgID)
			MessageEditor(update, QUserAge, keyboards.KUserAge)

		case "userSexBack":
			dbfellow.SetUserState("username", messageID, tgID)
			MessageEditor(update, QEnterUserName, keyboards.KUserName)

		//edit SEX

		case "editSex":
			DeleteMessage(chatID, messageID)
			MessageCompiler(chatID, QUserSex, keyboards.KEditUserSex)

		case "euМ", "euЖ", "euне указан":
			text := strings.TrimPrefix(data, "eu")
			dbfellow.UpdateTableWithText("users", "sex", text, tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		case "euserSexBack":
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		// AGE

		case "U18-24", "U25-34", "U35-44", "U45-54", "U55+":
			text := strings.TrimPrefix(data, "U")
			dbfellow.UpdateTableWithText("users", "age", text, tgID)
			MessageEditor(update, QInterests, keyboards.KInterests)

		case "userAgeBack":
			MessageEditor(update, QUserSex, keyboards.KUserSex)

		// edit AGE

		case "editAge":
			DeleteMessage(chatID, messageID)
			MessageCompiler(chatID, QUserAge, keyboards.KEditUserAge)

		case "eU18-24", "eU25-34", "eU35-44", "eU45-54", "eU55+":
			text := strings.TrimPrefix(data, "eU")
			dbfellow.UpdateTableWithText("users", "age", text, tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		case "euserAgeBack":
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		// INTERESTS

		case "культура", "психология", "экстрим", "природа", "ЗОЖ", "анти-ЗОЖ", "еда", "спорт", "путешествия", "гейминг":
			err := CheckboxCallback(update)
			if err != nil {
				log.Printf("Problem with checkboxes, user %s: %v", tgID, err)
			}

		case "saveInterests":
			choice := CollectMultiOptions(update)
			if SomethingChosen(choice) == true {
				dbfellow.SetUserState("aboutMe", messageID, tgID)
				dbfellow.UpdateTableWithText("users", "interests", choice, tgID)
				MessageEditor(update, QAboutMe, keyboards.KAboutMe)
			} else {
				PopUpInfoMessage(chatID, NothingWasChosen)
			}

		case "interestsBack":
			MessageEditor(update, QUserAge, keyboards.KUserAge)

		// edit INTERESTS

		case "editInterests":
			DeleteMessage(chatID, messageID)
			MessageCompiler(chatID, QInterests, keyboards.KEditInterests)

		case "esaveInterests":
			choice := CollectMultiOptions(update)
			if SomethingChosen(choice) == true {
				dbfellow.UpdateTableWithText("users", "interests", choice, tgID)
				BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)
			} else {
				PopUpInfoMessage(chatID, NothingWasChosen)
			}

		case "einterestsBack":
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		// ABOUT ME

		case "skipAboutMe":
			dbfellow.SetToDefault("users", "aboutMe", tgID)
			dbfellow.SetUserState("photo", messageID, tgID)
			MessageEditor(update, QPhoto, keyboards.KPhoto)

		case "aboutMeBack":
			dbfellow.SetUserStateToNeutral(tgID)
			MessageEditor(update, QInterests, keyboards.KInterests)

		// edit ABOUT ME

		case "editAboutMe":
			DeleteMessage(chatID, messageID)
			message, _ := MessageCompiler(chatID, QAboutMe, keyboards.KEditAboutMe)
			dbfellow.SetUserState("editAboutMe", message.MessageID, tgID)

		case "eaboutMeBack":
			dbfellow.SetUserStateToNeutral(tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		// PHOTO

		case "photoBack":
			dbfellow.SetUserState("aboutMe", messageID, tgID)
			MessageEditor(update, QAboutMe, keyboards.KAboutMe)

		// edit PHOTO

		case "editPhoto":
			DeleteMessage(chatID, messageID)
			message, _ := MessageCompiler(chatID, QPhoto, keyboards.KEditPhoto)
			dbfellow.SetUserState("editPhoto", message.MessageID, tgID)

		case "ephotoBack":
			dbfellow.SetUserStateToNeutral(tgID)
			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		// PREVIEW & SAVING

		case "preview", "previewAfterEditUser":

			user := dbfellow.GetUserData(tgID)

			text, photo := UserProfileTemplate(user, false, chatID)
			MessageWithPhoto(text, photo, keyboards.KEditOrSearch)

		case "editUserinfoOnPreview":

			BackToEditedUserinfoMessageFromCallback(update, keyboards.KEditUserinfo, tgID)

		case "go search":
			MessageCompiler(chatID, QPreShowInfo, keyboards.KSearch)

		case "search":
			err := Search(chatID, tgID)
			SendIfProblem(chatID, "Во время поиска произошла ошибка. Попробуем еще раз?", keyboards.KSearchRetry, err)

		//VIEWING SEARCH RESULTS

		case "nextCV":

			DeleteMessage(chatID, messageID)
			nextCV := dbfellow.GetNextIDString("users", "searchresults", "currentCV", tgID)
			if nextCV != "" {
				ShowCV(nextCV, chatID)
			} else {
				MessageCompiler(chatID, EverythingWasShown, keyboards.KEndOfFoundList)
			}

		case "watchCVsAgain":

			dbfellow.SetToDefault("users", "currentCV", tgID)
			firstCV := dbfellow.GetIDofTheFirstCVString("users", "searchResults", tgID)
			if firstCV != "" {
				ShowCV(firstCV, chatID)
			} else {
				MessageEditor(update, "Анкет для просмотра почему-то нет", keyboards.KDelete)
			}

		//LIKE

		case regLike.FindString(data):

			likedID := strings.TrimSuffix(data, "||like")
			iLiked, iWasLiked, likeID := dbfellow.AlreadyInLikes(likedID, tgID)

			if iWasLiked {
				DeleteMessage(chatID, messageID)
				keyboard := AddToOneButtonKeyboard(likeID, keyboards.SendMyUsername)
				MessageCompiler(chatID, ThisUserLikesYourTrip, keyboard)

			} else if iLiked {
				LD := dbfellow.GetLikeData(likeID)
				Like(update, LD.LikeToTgID, likeID, update.CallbackQuery.Message.From.UserName, tgID)

			} else {
				likeID := dbfellow.SaveLike(likedID, tgID)
				Like(update, likedID, likeID, update.CallbackQuery.Message.From.UserName, tgID)
			}

		//REQUEST

		case regSendMyUsername.FindString(data):

			userID := strings.TrimSuffix(data, "||sendMyUsername")
			SendUsername(chatID, userID, update.CallbackQuery.From.UserName, tgID)

		case regRepeatSendingUsername.FindString(data):

			userID := strings.TrimSuffix(data, "||repeatSendingUsername")
			SendUsername(chatID, userID, update.CallbackQuery.From.UserName, tgID)

		//MY REQUESTS

		case "showNextRequest":

			nextLikeID := dbfellow.GetNextIDString("users", "likeIDs", "currentR", tgID)
			DeleteMessage(chatID, messageID)

			if nextLikeID != "" {
				ShowRequest(chatID, nextLikeID)
			} else {
				DeleteMessage(chatID, messageID)
				dbfellow.SetToDefault("users", "requests", tgID)
				dbfellow.SetToDefault("users", "currentR", tgID)
				MessageCompiler(chatID, AllShown, keyboards.EmptyKeyboard)
			}

		//MY MATCHES

		case "showNextMatch":

			nextMatchID := dbfellow.GetNextIDString("users", "matches", "currentM", tgID)
			DeleteMessage(chatID, messageID)

			if nextMatchID != "" {
				err := ShowMatch(chatID, nextMatchID, tgID)
				SendIfProblem(chatID, SomthWrong, keyboards.EmptyKeyboard, err)
			} else {
				DeleteMessage(chatID, messageID)
				dbfellow.SetToDefault("users", "matches", tgID)
				dbfellow.SetToDefault("users", "currentM", tgID)
				MessageCompiler(chatID, AllShown, keyboards.EmptyKeyboard)
			}

		//FEEDBACK

		case regFeedback.FindString(data):

			chatID, _ := strconv.Atoi(strings.TrimSuffix(data, "||answerFeedback"))
			dbfellow.SetUserState("feedbackReply", chatID, tgID)
			MessageCompilerFeedbackReply(strings.TrimPrefix(update.CallbackQuery.Message.Text, "<b>💌 Новый фидбек: </b>\n"), int64(chatID))

		case "delete":
			DeleteMessage(chatID, messageID)
		}
	}
}
