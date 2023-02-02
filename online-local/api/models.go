package api

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Race struct {
	Id            string `json:"id"`
	OwnerId       string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	IsMultipleDay bool   `json:"is_multiple_day"`
}

type Competitor struct {
	Id             string `json:"id"`
	Identification string `json:"identification"`
	Name           string `json:"name"`
	Organization   string `json:"organization"`
	Region         string `json:"region"`
	Country        string `json:"country"`
	Birthday       string `json:"birthday"`
}

type Group struct {
	Id     string `json:"id"`
	RaceId string `json:"race_id"`
	Name   string `json:"name"`
}

type Course struct {
	Id     string         `json:"id"`
	RaceId string         `json:"race_id"`
	Name   string         `json:"name"`
	Chunks []*CourseChunk `json:"chunks"`
}

type CourseChunk struct {
	Code string `json:"code"`
}

type Result struct {
	Id           string         `json:"id"`
	Bib          string         `json:"bib"`
	CompetitorId string         `json:"competitor_id"`
	RaceId       string         `json:"race_id"`
	GroupId      string         `json:"group_id"`
	CourseId     string         `json:"course_id"`
	Qual         string         `json:"qual"`
	Chunks       []*ResultChunk `json:"chunks"`
}

type ResultChunk struct {
	Chip           string `json:"chip"`
	Identification string `json:"identification"`
	Code           string `json:"code"`
	Time           string `json:"time"`
}

type Chip struct {
	Name           string `json:"name"`
	Identification string `json:"identification"`
	RaceId         string `json:"race_id"`
	ResultId       string `json:"result_id"`
}
