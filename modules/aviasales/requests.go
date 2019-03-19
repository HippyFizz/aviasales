package aviasales

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestMethod string

var (
	MethodGet     RequestMethod = "GET"
	MethodHead    RequestMethod = "HEAD"
	MethodPost    RequestMethod = "POST"
	MethodPut     RequestMethod = "PUT"
	MethodPatch   RequestMethod = "PATCH" // RFC 5789
	MethodDelete  RequestMethod = "DELETE"
	MethodConnect RequestMethod = "CONNECT"
	MethodOptions RequestMethod = "OPTIONS"
	MethodTrace   RequestMethod = "TRACE"
)

func (s *Service) RequestPlaces(params url.Values, body io.Reader) ([]*WidgetFormat, error) {
	var places []PlacesResponse

	data, err := s.proxyRequest(PlacesServiceUrl, http.MethodGet, params, body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &places)
	if err != nil {
		return nil, err
	}

	rawResponse := make([]AviasalesResponse, len(places))
	for index := range places {
		rawResponse[index] = &places[index]
	}

	parsed, err := s.parseResponse(rawResponse)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (s *Service) proxyRequest(serviceUrl ServiceUrl, method RequestMethod, params url.Values, body io.Reader) (respBody []byte, err error) {
	client := &http.Client{}
	req, err := s.buildRequest(serviceUrl, method, params, body)

	if err != nil {
		return respBody, err
	}
	resp, err := client.Do(req)

	if err != nil {
		return respBody, err
	}

	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return respBody, err
	}

	return respBody, err
}

func (s *Service) buildRequest(serviceUrl ServiceUrl, method RequestMethod, params url.Values, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(string(method), fmt.Sprintf("%s%s", s.url, serviceUrl), body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, values := range params {
		for index := range values {
			q.Add(key, values[index])
		}
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (s *Service) parseResponse(response []AviasalesResponse) ([]*WidgetFormat, error) {
	var parsed []*WidgetFormat

	for index := range response {
		parsed = append(parsed, response[index].toWidgetFormat())
	}

	return parsed, nil
}
