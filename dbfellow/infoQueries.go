package dbfellow

import (
	"fmt"
	"log"

	"meet_bot/evil"
)

//gives a number of entries with diven tgID in a given DB table
func EntriesWithGivenTgID(tgID, table string) (int, error) {

	var count int

	str := fmt.Sprintf("select count(*) from %v where tgID = '%s'", table, tgID)

	err := Db.QueryRow(str).Scan(&count)
	if err != nil {
		evil.ProblemIn("EntriesWithGivenTgID", str, err)
		return count, err
	}

	return count, nil
}

//checks user's entry in table Users for "saved" flag
func UserProfileCompleted(tgID string) (saved bool) {

	err := Db.QueryRow("select saved from users where tgID = $1", tgID).Scan(&saved)
	if err != nil {
		log.Printf("Can't get info if user %s profile saved: %v", tgID, err)
	}
	return
}
