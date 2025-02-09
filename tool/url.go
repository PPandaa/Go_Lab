package tool

import (
	"GoLab/guard"
	"net/url"
)

func ParseStringToURL(target string) *url.URL {
	// Parse the URL
	target_url, err := url.Parse(target)

	if err != nil {
		guard.Logger.Fatal("Error parsing URL - " + err.Error())
	}

	return target_url
}
