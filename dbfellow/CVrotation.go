package dbfellow

//returns next ID from arrayColunm, increases counter in indexColumn
//returns 0 if next next index acceeds array length
func GetNextIDString(table, arrayColumn, indexColumn, tgID string) (nextID string) {

	index, arrayLength := GetIndex(table, arrayColumn, indexColumn, tgID)
	nextIndex := index + 1
	if IndexOk(nextIndex, arrayLength) {
		SetIndex(table, indexColumn, nextIndex, tgID)
		return GetStringValueByIndex(table, arrayColumn, nextIndex, tgID)
	}
	return
}

//gets value of the first element in array in DB.
//safe - checks if array has this element in advance. If it hasn't, returns 0
func GetIDofTheFirstCVString(table, arrayColumn, tgID string) string {

	length := GetArrayLength(table, arrayColumn, tgID)
	if IndexOk(1, length) {
		return GetStringValueByIndex(table, arrayColumn, 1, tgID)
	}
	return ""
}
