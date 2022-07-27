package dbfellow

import (
	"fmt"
	"log"
)

type UserData struct {
	TgID      string
	NameInBot string
	Age       string
	Sex       string

	Interests string
	AboutMe   string
	Photo     string

	Location string

	PartnersAge string
	PartnersSex string
}

//extracts userinfo from table Users and creates new UserData struct with it
func GetUserData(tgID string) UserData {

	var ud UserData

	query := fmt.Sprintf(`select nameInBot, age, sex, interests, aboutMe, photo, location, 
	partnersage, partnerssex from users where tgID = '%s'`, tgID)

	err := Db.QueryRow(query).Scan(&ud.NameInBot, &ud.Age, &ud.Sex, &ud.Interests, &ud.AboutMe,
		&ud.Photo, &ud.Location, &ud.PartnersAge, &ud.PartnersSex)
	if err != nil {
		log.Printf("Can't retrieve userinfo for user %s: %v", tgID, err)
	}

	return ud
}
