package aviasales

import "aviasales/modules/aviasales/config"

type Service struct {
	url     string
	Locales []Locale
	Types   []SearchType
}

type PlacesResponse struct {
	CityName     string            `json:"city_name"`
	Type         SearchType        `json:"type"`
	StateCode    string            `json:"state_code"`
	Cases        map[string]string `json:"cases"`
	CityCases    map[string]string `json:"city_cases"`
	CountryCases map[string]string `json:"country_cases"`
	Name         string            `json:"name"`
	CountryCode  string            `json:"country_code"`
	Code         string            `json:"code"`
	CityCode     string            `json:"city_code"`
	CountryName  string            `json:"country_name"`
	Coordinates  *Coordinate       `json:"coordinates"`
	Weight       int64             `json:"weight"`
	IndexStrings []string          `json:"index_strings"`
}

func (p *PlacesResponse) toWidgetFormat() *WidgetFormat {
	switch p.Type {
	case countrySearchType:
		return &WidgetFormat{Slug: p.Code, Title: p.Name, Subtitle: p.Name}
	case airportSearchType:
		return &WidgetFormat{Slug: p.Code, Title: p.Name, Subtitle: p.CityName}
	case citySearchType:
		return &WidgetFormat{Slug: p.Code, Title: p.Name, Subtitle: p.CountryName}
	default:
		return &WidgetFormat{Slug: p.Code, Title: "", Subtitle: ""}
	}
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

var aviasalesApi *Service

func newApi(config *config.Config) error {
	aviasalesApi = &Service{url: config.Url, Locales: locales, Types: SearchTypes}
	return nil
}

func Manager(config *config.Config) *Service {
	if err := newApi(config); err != nil {
		panic(err)
	}
	return aviasalesApi
}
