package utils

import (
	"compress/gzip"
	"io"
	"net/http"

	"github.com/mozillazg/request"
)

func get(url string) (reader io.Reader, err error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get(url)
	if err != nil {
		return
	}
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
	} else {
		reader = resp.Body
	}

	return
}
