package structures

const ElementAfficher = 2

// --------------------------------------------------------------------
var IndexRelations INDEX_Relations

type Relations struct {
	Id                int                 `json:"id"`
	DatesLocalisation map[string][]string `json:"datesLocations"`
}

type INDEX_Relations struct {
	Index []Relations `json:"index"`
}

// --------------------------------------------------------------------
var IndexLocations INDEX_Locations

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type INDEX_Locations struct {
	Index []Location `json:"index"`
}

// --------------------------------------------------------------------

var IndexDates INDEX_Dates

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type INDEX_Dates struct {
	Index []Dates `json:"index"`
}

// --------------------------------------------------------------------

var Apiartistes []Artiste

type Artiste struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationdate"`
	FirstAlbum   string   `json:"firstalbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertdates"`
	Relations    string   `json:"relations"`
}

// --------------------------------------------------------------------

var API_Globale []ApiGlobale

type ApiGlobale struct {
	Id_ARTISTE           int
	Image_ARTISTE        string
	Name_ARTISTE         string
	Members_ARTISTE      []string
	CreationDate_ARTISTE int
	FirstAlbum_ARTISTE   string
	Locations_ARTISTE    string
	ConcertDates_ARTISTE string
	Relations_ARTISTE    string
}
