package models

type Error struct {
	ErrorCode    int
	ErrorMessage string
}
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type SingleArtist struct {
	Artist   Artist
	Relation Relation
	Location Locations
	Dates    Date
}

type Relation struct {
	Id           int                 `json:"id"`
	DateLocation map[string][]string `json:"datesLocations"`
}
type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Filters struct {
	NumberOfMembers   []int
	CreationDate      []int
	FirstAlbumRelease []int
	Location          string
	Shearch           string
}

type LocationType struct {
	Index []Locations `json:"index"`
}
type RelationType struct {
	Index []Relation `json:"index"`
}

type DatesType struct {
	Index []Date `json:"index"`
}

type FetchedType interface {
	LocationType | RelationType | DatesType | Locations | Relation | Artist | Date
}
