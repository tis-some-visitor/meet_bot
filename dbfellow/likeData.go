package dbfellow

import (
	"fmt"
	"log"

	"meet_bot/evil"
)

type LikeData struct {
	LikeID       int
	LikeFromTgID string
	LikeToTgID   string
	Match        bool
}

func GetLikeData(likeID int) (ld LikeData) {

	err := Db.QueryRow(`select likeID, likeFromTgID, likeToTgID, match from likes 
	where likeID = $1`, likeID).Scan(&ld.LikeID, &ld.LikeFromTgID, &ld.LikeToTgID, &ld.Match)
	if err != nil {
		log.Printf("Can't get info about like №%v: %v", likeID, err)
	}
	return
}

//given a tgID finds chatID
func GetUsernameAndChatID(tgID string) (chatID int64, username string) {

	err := Db.QueryRow("select num, title from states_ where owl = $1", tgID).Scan(&chatID, &username)
	if err != nil {
		log.Printf("Problem in GetUsernameAndChatID, tgID %s, error %v", tgID, err)
	}
	return
}

//sets "match" = true in table Likes, saves username that was sent
func Match(likeID int, tgID, username string) {

	SetMatchTrue(likeID)
	SaveUsername(tgID, username)

}

func SetMatchTrue(likeID int) {

	query := fmt.Sprintf("update likes set match = true where likeID = %v", likeID)

	_, err := Db.Exec(query)
	if err != nil {
		evil.ProblemIn("SetMatchTrue", query, err)
	}
}

func SaveUsername(tgID, username string) {

	query := fmt.Sprintf("update states_ set title = '%s' where owl = '%s'", username, tgID)

	_, err := Db.Exec(query)
	if err != nil {
		evil.ProblemIn("SaveUsername", query, err)
	}
}

//create new entry in Likes with 2 tgIDs. Returns likeID
func SaveLike(likedID, tgID string) int {

	str := fmt.Sprintf("insert into likes (likeFromTgID, likeToTgID) values ('%s', '%s')", tgID, likedID)
	_, err := Db.Exec(str)
	if err != nil {
		evil.ProblemIn("SaveLike", str, err)
		return 0
	}

	query := fmt.Sprintf("select likeID from likes where likeFromTgID = '%s' and likeToTgID = %v", tgID, likedID)

	var likeID int
	err = Db.QueryRow(query).Scan(&likeID)
	if err != nil {
		evil.ProblemIn("SaveLike", query, err)
		return 0
	}

	return likeID
}

func AlreadyInLikes(likedID, tgID string) (iLiked, iWasLiked bool, likeID int) {

	query := fmt.Sprintf("select likeID from likes where likeFromTgID = %v and likedToTgID = %v", tgID, likedID)

	var iLikedInt int
	err := Db.QueryRow(query).Scan(&iLikedInt)
	if err != nil {
		evil.ProblemIn("AlreadyInLikes", query, err)
	}

	query = fmt.Sprintf("select likeID from likes where likeFromTgID = %v and likedToTgID = %v", likedID, tgID)

	var iWasLikedInt int
	err = Db.QueryRow(query).Scan(&iWasLikedInt)
	if err != nil {
		evil.ProblemIn("AlreadyInLikes", query, err)
	}

	if iWasLikedInt > 0 {
		return false, true, iWasLikedInt
	} else if iLikedInt > 0 {
		return true, false, iLikedInt
	} else {
		return false, false, 0
	}

}
