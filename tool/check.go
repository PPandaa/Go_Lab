package tool

import (
	"net/http"
)

func IsSiteReachable(url string) bool {

	http_client := &http.Client{}
	site_status := false

	request, _ := http.NewRequest("GET", url, nil)
	response, _ := http_client.Do(request)
	if response.StatusCode == 200 {
		site_status = true
	}

	return site_status

}
