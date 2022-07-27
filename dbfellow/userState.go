package dbfellow

import (
	"fmt"
	"log"
)

//return user's state and latest keyboard message number
func GetUserState(tgID string) (string, int) {

	var state string
	var mesNO int

	err := Db.QueryRow("select stage, mesNo from users where tgID = $1", tgID).Scan(&state, &mesNO)

	if err != nil {
		log.Printf("Can't get state and mesNo of user %s: %v", tgID, err)
	}

	return state, mesNO
}

//set user's state for a given one, writes the current keyboard message number to the column mesNo
func SetUserState(state string, messageID int, tgID string) error {

	str := fmt.Sprintf("update users set (stage, mesNo) = ('%s', %v) where tgID = '%s'", state, messageID, tgID)
	_, err := Db.Exec(str)
	if err != nil {

		err = RetryExec(str)
		if err != nil {
			err = fmt.Errorf("Can't set user state to %s for user %s: %w", state, tgID, err)
			log.Print(err)
			return err
		}
	}
	return nil
}

//set user's state to empty string, mesNo to 0
func SetUserStateToNeutral(tgID string) {

	_, err := Db.Exec("update users set (stage, mesNO) = (default, default) where tgID = $1", tgID)
	if err != nil {
		err = fmt.Errorf("Can't set user state to neutral for user %s: %w", tgID, err)
		log.Print(err)
	}
}
