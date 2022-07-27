package dbfellow

import (
	"fmt"
	"log"
)

func GetMyRequests(tgID string) []string {

	query := fmt.Sprintf("select likeFromTgID from likes where likeToTgID = '%s' and match = false order by liketime desc;", tgID)
	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error during query, getting requests for user %s: %v", tgID, err)
	}

	defer rows.Close()

	var IDs []string

	for rows.Next() {

		var ID string
		if err = rows.Scan(&ID); err != nil {
			log.Printf("Scan error in cycle while getting requests for user %s: %v", tgID, err)
		}
		IDs = append(IDs, ID)

	}

	if err := rows.Err(); err != nil {
		log.Printf("Scan error while getting requests for user %s: %v", tgID, err)
	}

	return IDs
}

func GetMyMatches(tgID string) []string {

	query := fmt.Sprintf("select likeFromTgID, likeToTgID from likes where (likeToTgID = '%s' or likeFromTgID = '%s') and match = true order by liketime desc;", tgID, tgID)
	rows, err := Db.Query(query)
	if err != nil {
		log.Printf("Error during query, getting matches for user %s: %v", tgID, err)
	}

	defer rows.Close()

	var IDs []string

	for rows.Next() {

		var ID1, ID2 string
		if err = rows.Scan(&ID1, &ID2); err != nil {
			log.Printf("Scan error in cycle while getting matches for user %s: %v", tgID, err)
		}
		if ID1 == tgID {
			IDs = append(IDs, ID2)
		} else {
			IDs = append(IDs, ID1)
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Scan error while getting matches for user %s: %v", tgID, err)
	}

	return IDs
}
