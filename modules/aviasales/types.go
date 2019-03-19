package aviasales

type SearchType string

type AviasalesResponse interface {
	toWidgetFormat() *WidgetFormat
}

type WidgetFormat struct {
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

var (
	citySearchType    SearchType = "city"
	airportSearchType SearchType = "airport"
	countrySearchType SearchType = "country"
	SearchTypes                  = []SearchType{citySearchType, airportSearchType, countrySearchType}
)
