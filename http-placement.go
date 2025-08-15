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

func (p *HttpPlacement) Extract(r *http.Request) (data string, err error) {
	if len(p.Header) != 0 {
		return r.Header.Get(p.Header), nil
	}
	if len(p.PathParam) != 0 {
		return r.PathValue(p.PathParam), nil
	}
	if len(p.UrlParam) != 0 {
		return r.URL.Query().Get(p.UrlParam), nil
	}
	if p.Body {
		rawData, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		return string(rawData), err
	}
	return "", ErrPlacementUnspecified
}
