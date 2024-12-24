package api

// Sert pour affiché un nombre spécifique d'ID
const ElementsAffichages = 2

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| STRUCTURE ARTISTES |------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

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

var ApiArtistes Artiste

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| STRUCTURE DATES |---------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type ApiDATES struct {
	Index []Dates `json:"index"`
}

var ApiDates ApiDATES

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| STRUCTURE LOCATIONS |-----------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Api_Locations struct {
	Index []Location `json:"index"`
}

var ApiLocations Api_Locations

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| STRUCTURE RELATIONS |-----------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

type Relations struct {
	Id                int                 `json:"id"`
	DatesLocalisation map[string][]string `json:"datesLocations"`
}

type Api_Relations struct {
	Index []Relations `json:"index"`
}

var ApiRelations Api_Relations

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| STRUCTURE TRIAGE |--------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

type Tri_Fonctions struct {
	// variales apres le reformatage et nettoyage
}

var TriFonctions Tri_Fonctions

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------| STRUCTURE GLOBALE |-------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------

type Globale_Structure_Donnees struct {
	ID_ARTISTE            int
	IMAGE                 string
	NAME_ARTISTE          string
	MEMBERS_ARTISTE       []string
	CREATION_DATE_ARTISTE int
	FIRST_ALBUM_ARTISTE   string
	LOCATIONS_ARTISTE     string
	CONCERT_DATES_ARTISTE string
	RELATIONS_ARTISTE     string
	//----------------------------------------
	INDEX_DATES int
	ID_DATES    int
	DATES_DATES []string
	//----------------------------------------
	INDEX_LOCATIONS     int
	ID_LOCATIONS        int
	LOCATIONS_LOCATIONS []string
	DATES_LOCATIONS     []string
	//---------------------------------------
	INDEX_RELATIONS          int
	ID_RELATIONS             int
	DATESLOCATIONS_RELATIONS map[string][]string
	DATES_RELATIONS          []string
	LOCATIONS_RELATIONS      []string
}

var GlobaleStructure Globale_Structure_Donnees

// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------------