package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	geocodingURL = "http://maps.googleapis.com/maps/api/geocode/json"
)

type GeocodingResult struct {
	Results []*GeocodingAddress `json:"results,omitempty"`
}

type GeocodingAddress struct {
	FormattedAddress string            `json:"formatted_address,omitempty"`
	Geometry         GeocodingGeometry `json:"geometry,omitempty"`
}

type GeocodingGeometry struct {
	Location GeometryLocation `json:"location,omitempty"`
}

type GeometryLocation struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

func (l *GeometryLocation) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Lng)
}

func Geocoding(location string, lat, lng float64) (*GeocodingResult, error) {
	qs, err := url.ParseQuery("")
	Check(err)
	if location != "" {
		qs.Add("address", location)
	}
	if lat != 0 {
		qs.Add("lat", strconv.FormatFloat(lat, 'f', 6, 64))
	}
	if lng != 0 {
		qs.Add("lng", strconv.FormatFloat(lng, 'f', 6, 64))
	}
	qs.Add("sensor", "false")
	u := geocodingURL + "?" + qs.Encode()

	resp, err := http.Get(u)
	Check(err)

	r := new(GeocodingResult)
	err = json.NewDecoder(resp.Body).Decode(r)
	return r, err
}
