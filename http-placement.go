package profiles

import (
	"errors"
	"io"
	"net/http"
)

var (
	ErrPlacementUnspecified = errors.New("placement of data is unspecified")
)

type HttpPlacement struct {
	Header    string `json:"header"`
	UrlParam  string `json:"urlParam"`
	PathParam string `json:"pathParam"`
	Body      bool   `json:"body"`
}

func Extract(placement *HttpPlacement, r *http.Request) (data string, err error) {
	if len(placement.Header) != 0 {
		return r.Header.Get(placement.Header), nil
	}
	if len(placement.PathParam) != 0 {
		return r.PathValue(placement.PathParam), nil
	}
	if len(placement.UrlParam) != 0 {
		return r.URL.Query().Get(placement.UrlParam), nil
	}
	if placement.Body {
		rawData, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		return string(rawData), err
	}
	return "", ErrPlacementUnspecified
}
