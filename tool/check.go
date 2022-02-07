package tool

import (
	"net/http"
	"strings"
)

func IsSiteReachable(url string) bool {

	http_client := &http.Client{}
	site_status := false

	var rURL string
	if strings.HasSuffix(url, "/graphql") {
		temp := strings.Split(url, "/graphql")
		rURL = temp[0]
	} else {
		rURL = url
	}

	request, _ := http.NewRequest("GET", rURL, nil)
	response, _ := http_client.Do(request)
	if response.StatusCode == 200 {
		site_status = true
	}

	return site_status

}

func IsStringDuplicate(target string, elements []string) bool {

	result := false
	for _, elelement := range elements {
		if elelement == target {
			result = true
		}
	}
	return result

}
