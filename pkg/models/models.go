package models

// CharacterDataWrapper Object
type CharacterDataWrapper struct {
	Code            int                    `json:"code"`
	Status          string                 `json:"status"`
	Copyright       string                 `json:"copyright"`
	AttributionText string                 `json:"attributionText"`
	AttributionHTML string                 `json:"attributionHTML"`
	Data            CharacterDataContainer `json:"data"`
	Etag            string                 `json:"etag"`
}

// ComicDataWrapper Object
type ComicDataWrapper struct {
	Code            int                `json:"code"`
	Status          string             `json:"status"`
	Copyright       string             `json:"copyright"`
	AttributionText string             `json:"attributionText"`
	AttributionHTML string             `json:"attributionHTML"`
	Data            ComicDataContainer `json:"data"`
	Etag            string             `json:"etag"`
}

// CharacterDataContainer Object
type CharacterDataContainer struct {
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Count   int         `json:"count"`
	Results []Character `json:"results"`
}

// ComicDataContainer Object
type ComicDataContainer struct {
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Comic `json:"results"`
}

// Character Object
type Character struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Modified    string     `json:"modified"` // Date
	ResourceURI string     `json:"resourceURI"`
	Urls        []URL      `json:"urls"`
	Thumbnail   Image      `json:"thumbnail"`
	Comics      ComicList  `json:"comics"`
	Stories     StoryList  `json:"stories"`
	Events      EventList  `json:"events"`
	Series      SeriesList `json:"series"`
}

// Comic Object
type Comic struct {
	ID                 int            `json:"id"`
	DigitalID          int            `json:"digitalId"`
	Title              string         `json:"title"`
	IssueNumber        float64        `json:"issueNumber"`
	VariantDescription string         `json:"variantDescription"`
	Description        string         `json:"description"`
	Modified           string         `json:"modified"` //Date
	Isbn               string         `json:"isbn"`
	Upc                string         `json:"upc"`
	DiamondCode        string         `json:"diamondCode"`
	Ean                string         `json:"ean"`
	Issn               string         `json:"issn"`
	Format             string         `json:"format"`
	PageCount          int            `json:"pageCount"`
	TextObjects        []TextObject   `json:"textObjects"`
	ResourceURI        string         `json:"resourceURI"`
	Urls               []URL          `json:"urls"`
	Series             SeriesSummary  `json:"series"`
	Variants           []ComicSummary `json:"variants"`
	Collections        []ComicSummary `json:"collections"`
	CollectedIssues    []ComicSummary `json:"collectedIssues"`
	Dates              []ComicDate    `json:"dates"`
	Prices             []ComicPrice   `json:"prices"`
	Thumbnail          Image          `json:"thumbnail"`
	Images             []Image        `json:"images"`
	Creators           CreatorList    `json:"creators"`
	Characters         CharacterList  `json:"characters"`
	Stories            StoryList      `json:"stories"`
	Events             EventList      `json:"events"`
}

// TextObject Object
type TextObject struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

// URL Object
type URL struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

// ComicDate Object
type ComicDate struct {
	Type string `json:"type"`
	Date string `json:"date"` // Date
}

// ComicPrice Object
type ComicPrice struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"` // Date
}

// Image Object
type Image struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

// ComicList Object
type ComicList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []ComicSummary `json:"items"`
}

// ComicSummary Object
type ComicSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

// CreatorList Object
type CreatorList struct {
	Available     int              `json:"available"`
	Returned      int              `json:"returned"`
	CollectionURI string           `json:"collectionURI"`
	Items         []CreatorSummary `json:"items"`
}

// CreatorSummary Object
type CreatorSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

// CharacterList Object
type CharacterList struct {
	Available     int                `json:"available"`
	Returned      int                `json:"returned"`
	CollectionURI string             `json:"collectionURI"`
	Items         []CharacterSummary `json:"items"`
}

// CharacterSummary Object
type CharacterSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

// SeriesList Object
type SeriesList struct {
	Available     int             `json:"available"`
	Returned      int             `json:"returned"`
	CollectionURI string          `json:"collectionURI"`
	Items         []SeriesSummary `json:"items"`
}

// SeriesSummary Object
type SeriesSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

// StoryList Object
type StoryList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []StorySummary `json:"items"`
}

// StorySummary Object
type StorySummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

// EventList Object
type EventList struct {
	Available     int            `json:"available"`
	Returned      int            `json:"returned"`
	CollectionURI string         `json:"collectionURI"`
	Items         []EventSummary `json:"items"`
}

// EventSummary Object
type EventSummary struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

// API MarvelCharacterResponse Object
type MarvelCharacterResponse struct {
	ResultsCount int               `json:"count"`
	Characters   []MarvelCharacter `json:"characters"`
}

// API MarvelCharacter Object
type MarvelCharacter struct {
	Name        string `json:"name"`
	CharacterID int    `json:"characterID"`
}

// API MarvelComicResponse Object
type MarvelComicResponse struct {
	ResultsCount int           `json:"count"`
	Comics       []MarvelComic `json:"comics"`
}

// API MarvelComic Object
type MarvelComic struct {
	Title      string   `json:"title"`
	Characters []string `json:"characters"`
}
