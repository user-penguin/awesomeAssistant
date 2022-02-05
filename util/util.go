package util

import (
	"log"
	"net/url"
)

func UrlToCanonical(base string, path string, values map[string]string) (error, string) {
	Url, err := url.Parse(base)
	if err != nil {
		log.Fatal(err.Error())
		return err, ""
	}
	Url.Path += path
	params := url.Values{}
	for key, value := range values {
		params.Add(key, value)
	}
	Url.RawQuery = params.Encode()
	return nil, Url.String()
}
