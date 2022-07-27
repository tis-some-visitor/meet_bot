package dbfellow

import (
	"fmt"
	"log"
	"meet_bot/evil"

	"github.com/lib/pq"
)

//saves callback data to a specified table & column.
//converts data to string with fmtSprint for safety
func UpdateTableWithText(table, column, text interface{}, tgID string) {

	datastring := fmt.Sprintf("%v", text)

	str := fmt.Sprintf("update %s set %s = '%s' where tgID = '%s'", table, column, datastring, tgID)

	_, err := Db.Exec(str)
	if err != nil {
		log.Printf("Can't execute update %s: %v", str, err)
	}
}

//saves a slice if integers as an SQL array
func UpdateTableWithStringArray(table, column, tgID string, array []string) error {

	query := fmt.Sprintf("update %s set %s = $1 where tgID = '%s'", table, column, tgID)

	_, err := Db.Exec(query, pq.Array(array))
	if err != nil {
		err := fmt.Errorf("Can't save string array with query %s: %v", query, err)
		log.Print(err)
	}
	return err
}

//sets a given cell to default
func SetToDefault(table, column, tgID string) {

	str := fmt.Sprintf("update %s set %s = default where tgID = '%s'", table, column, tgID)
	_, err := Db.Exec(str)
	if err != nil {
		log.Printf("Can't execute update %s: %v", str, err)
	}
}

//gets data (of type string) from a given cell
func GetSavedData(table, column, tgID string) string {

	var values string

	query := fmt.Sprintf("select %s from %s where tgID = '%s'", column, table, tgID)

	err := Db.QueryRow(query).Scan(&values)
	if err != nil {
		log.Printf("Can't get saved data (%s): %v", query, err)
	}

	return values
}

//sets "saved" switch in table Users (true or false)
func SavedSwitch(value bool, tgID string) {

	query := fmt.Sprintf("update users set saved = %v where tgID = '%s'", value, tgID)
	_, err := Db.Exec(query)
	if err != nil {
		evil.ProblemIn("SavedSwitch", query, err)
	}
}

func DropEntry(table, column string, ID interface{}) {

	query := fmt.Sprintf("delete from %s where %s = %v", table, column, ID)
	_, err := Db.Exec(query)
	if err != nil {
		evil.ProblemIn("DropEntry", query, err)
	}
}
