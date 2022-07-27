package dbfellow

import (
	"fmt"
	"log"

	"meet_bot/evil"
)

func GetIndex(table, arrayColumn, indexColumn, tgID string) (index, arrayLength int) {

	str := fmt.Sprintf("select %s, cardinality(%s) from %s where tgID = '%s'", indexColumn, arrayColumn, table, tgID)
	err := Db.QueryRow(str).Scan(&index, &arrayLength)
	if err != nil {
		log.Printf("Problem in GetIndex; query %s; error %v", str, err)
		return 0, 0
	}
	return
}

func GetArrayLength(table, arrayColumn, tgID string) (arrayLength int) {

	str := fmt.Sprintf("select cardinality(%s) from %s where tgID = '%s'", arrayColumn, table, tgID)
	err := Db.QueryRow(str).Scan(&arrayLength)
	if err != nil {
		log.Printf("Problem in GetIndex; query %s; error %v", str, err)
		return 0
	}
	return
}

func SetIndex(table, indexColumn string, index int, tgID string) {

	str := fmt.Sprintf("update %s set %s = %d where tgID = '%s'", table, indexColumn, index, tgID)
	_, err := Db.Exec(str)
	if err != nil {
		log.Printf("Problem in SetIndex; query %s; error %v", str, err)
	}
}

func IndexOk(index, arrayLength int) bool {

	if index > 0 && index <= arrayLength {
		return true
	} else {
		return false
	}
}

func GetStringValueByIndex(table, arrayColumn string, index int, tgID string) (nextID string) {

	str := fmt.Sprintf("select %s[%v] from %s where tgID = '%s'", arrayColumn, index, table, tgID)
	err := Db.QueryRow(str).Scan(&nextID)
	if err != nil {
		log.Printf("Problem selecting trip by index in GetNextTripID; query %s; error %v", str, err)
	}
	return
}

//deletes given ID from array in given column containing array. Substructs 1 from relevant index column.
func DeleteStringFromArray(table, arrayColumn string, ID string, tgID string) {

	query := fmt.Sprintf("update %s set %s = array_remove(%s, %v) where tgID = '%s'", table, arrayColumn, arrayColumn, ID, tgID)
	_, err := Db.Exec(query)
	if err != nil {
		evil.ProblemIn("DeleteFromArray", query, err)
	}

}

func SubstructOneFromIndexColumn(table, indexColumn, tgID string) {

	query := fmt.Sprintf("update %s set %s = %s - 1 where tgID = '%s' and %s > 1", table, indexColumn, indexColumn, tgID, indexColumn)
	_, err := Db.Exec(query)
	if err != nil {
		evil.ProblemIn("SubstructOneFromIndexColumn", query, err)
	}
}
