package api

/* Variables */
const ElementAfficher = 2

// --------------------------------------------------------------------------------------------------------------------------------- //
// Gerer l'index pour l'import d'api

var IndexLocations Index_Locations // ----------- | INDEX API LOCATIONS | ----------- |

type Index_Locations struct {
	Locations []Location `json:"index"`
}

var IndexDates Index_Dates // ------------------- | INDEX API DATES | --------------- |

type Index_Dates struct {
	Dates []Dates `json:"index"`
}

var IndexRelations Index_Relations // ----------- | INDEX API RELATIONS | ----------- |

type Index_Relations struct {
	Relations []Relations `json:"index"`
}

// --------------------------------------------------------------------------------------------------------------------------------- //
// DONNEES DE L'API DES ARTISTES

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

// --------------------------------------------------------------------------------------------------------------------------------- //
// DONNEES DE L'API DESRELATIONS

type Relations struct {
	Id                int                 `json:"id"`
	DatesLocalisation map[string][]string `json:"datesLocations"`
}

// --------------------------------------------------------------------------------------------------------------------------------- //
// DONNEES DE L'API DES LIEUX

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// --------------------------------------------------------------------------------------------------------------------------------- //
// DONNEES DE L'API DES DATES

type Dates struct {
	Id        int      `json:"id"`
	Tab_Dates []string `json:"dates"`
}

// --------------------------------------------------------------------------------------------------------------------------------- //
// DONNEES APRES LES SORTIES DES FONCTIONS DE TRIES ET GESTIONS
var VarFonctions Var_Fonctions

// Structure globale des données après modification

type Var_Fonctions struct {
	// Map des dates formatées par ID
	DATE_FINALE map[int][]string

	// Map des lieux formatés par ID
	LOCATIONS_FINALES map[int][]string
}

// --------------------------------------------------------------------------------------------------------------------------------- //
// STRUCTURE GLOBALE ENVELOPPANT TOUTES LES VARIABLES A AFFICHEEES SUR LE SITE

type API_Globale struct {
	INDEX_LIEUX     Index_Locations
	INDEX_RELATIONS Index_Relations
	INDEX_DATES     Index_Dates

	DATES_FORMATEES Var_Fonctions

	// ---------------------------------------------- //
	// VARIABLE A MODIFIER EN FONCTIONS DU CODE FINAL //
	// ---------------------------------------------- //

	// ID_ARTISTE    Artiste
	// IMAGE         string
	// NOM           string
	// MEMBRE        []string
	// DATE_CREATION int
	// PREMIER_ALBUM string
	// LIEUX         string
	// DATE_CONCERT  string
	// RELATION      string

	// ID_DATES Dates
	// DATES    []string

	// ID_LOCATIONS    Location
	// LOCATIONS       []string
	// DATES_LOCATIONS string

	// ID_RELATIONS    Relations
	// DATES_RELATIONS map[string][]string
}
