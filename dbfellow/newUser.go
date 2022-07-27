package dbfellow

import "fmt"

//creates a new entry in table Users, using user's hashed ID, and firstname.
//if insert returns an error, retries 5 times with exponential backoff
func CreateNewUser(tgID, firstName string) error {

	query := fmt.Sprintf("insert into users (tgID, firstname) values('%s', '%s') on conflict (tgid) do update set (stage, mesno, searchresults, currentcv) = (default, default, default, default)", tgID, firstName)
	_, err := Db.Exec(query)

	if err != nil {
		err = RetryExec(query)
		return err
	}

	return nil
}

//creates a new entry in table States, using user's ID and its hashed version.
//does nothing on conflict (if such entry already exists).
//if insert returns an error, retries 5 times with exponential backoff
func CreateNewState(tgID string, ID int64) error {

	query := fmt.Sprintf("insert into states_ (owl, num) values('%s', %v) on conflict do nothing", tgID, ID)
	_, err := Db.Exec(query)
	if err != nil {
		err = RetryExec(query)
		return err
	}

	return nil
}
