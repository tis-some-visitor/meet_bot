package dbfellow

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type SearchResult struct {
	TgID      string
	Age       string
	Sex       string
	Interests string
	Location  string
	Rate      int64
}

//takes UserData, searches all entries in users view (in database) to find all entries
//appropriate by mutual choices of partner's sex and age. Example: if user is a 18-24 years old male and looking for female
//no matter what age, function will choose all females who are looking for 18-24 old male or doesn't care for sex and age
//of the partner.
//function parses each found entry into a new SearchResult struct, and returns a slice of them.
func SelectFromDBbySexAndAge(tgID string, ud UserData) []SearchResult {

	var query string

	if ud.PartnersSex == "не важно" && ud.PartnersAge == "не важно" {

		query = fmt.Sprintf(`select * from (select * from users where tgID != '%s') a 
		where (a.PartnersSex = 'не важно' or a.PartnersSex = '%s') and 
		(a.PartnersAge = 'не важно' or a.PartnersAge like '%%%s%%')`, tgID, ud.Sex, ud.Age)

	} else if ud.PartnersSex == "не важно" && ud.PartnersAge != "не важно" {

		query = fmt.Sprintf(`select * from (select * from (select * from users where tgID != '%s') a 
		where (a.PartnersSex = 'не важно' or a.PartnersSex = '%s') and 
		(a.PartnersAge = 'не важно' or a.PartnersAge like '%%%s%%')) b
		where b.Age like '%%%s%%'`, tgID, ud.Sex, ud.Age, ud.PartnersAge)

	} else if ud.PartnersSex != "не важно" && ud.PartnersAge == "не важно" {

		query = fmt.Sprintf(`select * from (select * from (select * from users where tgID != '%s') a 
		where (a.PartnersSex = 'не важно' or a.PartnersSex = '%s') and 
		(a.PartnersAge = 'не важно' or a.PartnersAge like '%%%s%%')) b
		where b.Sex = '%s'`, tgID, ud.Sex, ud.Age, ud.PartnersSex)

	} else if ud.PartnersSex != "не важно" && ud.PartnersAge != "не важно" {

		query = fmt.Sprintf(`select * from (select * from (select * from users where tgID != '%s') a 
		where (a.PartnersSex = 'не важно' or a.PartnersSex = '%s') and 
		(a.PartnersAge = 'не важно' or a.PartnersAge like '%%%s%%')) b
		where b.sex = '%s' and b.age like '%%%s%%'`, tgID, ud.Sex, ud.Age, ud.PartnersSex, ud.PartnersAge)
	}

	rows, err := Db.Query(query)

	if err != nil {
		log.Printf("Error during search query (%s): %v", query, err)
	}
	defer rows.Close()

	found := make([]SearchResult, 0)

	for rows.Next() {

		var sr SearchResult

		if err := rows.Scan(&sr.TgID, &sr.Age, &sr.Sex, &sr.Interests, &sr.Location); err != nil {
			log.Printf("Error in scan during search query (%s): %v", query, err)
		}

		found = append(found, sr)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error in scan during search query (%s): %v", query, err)
	}

	return found
}

//searches in table Likes for ID's that have already been liked by user or have already liked him
func FindLikesFrom(tgID string) []string {

	rows, err := Db.Query("select likeToTgID from likes where likeFromTgID = $1", tgID)
	if err != nil {
		err = fmt.Errorf("Error in FindLikesrom for %v: %w", tgID, err)
		log.Print(err)
	}

	defer rows.Close()

	AlreadyLiked := []string{}

	for rows.Next() {
		liked := ""

		if err := rows.Scan(&liked); err != nil {
			err = fmt.Errorf("Error during scan in FindLikesFrom %v: %w", tgID, err)
			log.Print(err)
		}

		AlreadyLiked = append(AlreadyLiked, liked)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during scan in FindLikesFrom %v: %v", tgID, err)
	}

	return AlreadyLiked
}

func ExcludeLikes(liked []string, sr []SearchResult) (cleanedResults []SearchResult) {

	for _, result := range sr {

		if !IsStringInSlice(liked, result.TgID) {

			cleanedResults = append(cleanedResults, result)
		}

	}
	return
}

func IsStringInSlice(slice []string, item string) bool {

	for _, elm := range slice {
		if elm == item {
			return true
		}
	}
	return false
}

//compairs each SearchResult struct to a user's data, adding score to SearchResult's rating.
//returns slice of SearchResult with Rate field filled.
func RatingCalculation(ud UserData, sr []SearchResult) []SearchResult {

	var ResultsWithRate []SearchResult

	for _, r := range sr {

		location := LocationMatch(ud, r)
		interests := InterestsMatch(ud, r)

		r.Rate = location + interests

		if r.Rate > 0 {
			ResultsWithRate = append(ResultsWithRate, r)
		}
	}

	return ResultsWithRate
}

//sorts rated SearchResults from most high rated to least and forms a slice  in rated order.
func SearchResultSorter(sr []SearchResult) []string {

	sort.Slice(sr, func(i, j int) bool { return sr[i].Rate > sr[j].Rate })

	var IDs []string

	for _, result := range sr {

		IDs = append(IDs, result.TgID)
	}

	return IDs

}

func LocationMatch(ud UserData, r SearchResult) int64 {

	var rate int64

	if r.Location == ud.Location {

		rate += 20

	}
	return rate
}

func FindSameInStringSlices(uInt, pInt string) (same []string) {

	uInterests := strings.Split(uInt, ", ")
	pInterests := strings.Split(pInt, ", ")

	for _, interest := range uInterests {

		for _, interest2 := range pInterests {

			if interest == interest2 {

				same = append(same, interest)
			}
		}
	}
	return
}

func InterestsMatch(ud UserData, sr SearchResult) int64 {

	same := FindSameInStringSlices(ud.Interests, sr.Interests)
	return int64(len(same))
}
