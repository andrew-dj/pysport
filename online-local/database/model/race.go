package model

import (
	"database/sql"
)

type Race struct {
	Id          string  `json:"id"`
	OwnerId     *string `json:"owner_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}

func GetRaces(db *sql.DB) []*Race {
	var races []*Race
	results, err := db.Query("SELECT * FROM race")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var race Race
		err = results.Scan(&race.Id, &race.OwnerId, &race.Title, &race.Description)
		if err != nil {
			panic(err.Error())
		}
		races = append(races, &race)
	}
	return races
}