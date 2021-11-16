package auth

import (
	"GoLab/dependency"
	"GoLab/server"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

var (
	IFPToken string
)

func CloudIFPToken() {

	for {
		timestamp := time.Now()
		options := &newSRPTokenOptions{Timestamp: &timestamp}
		result := newSrpToken("OEE", options)
		httpClient := &http.Client{}
		request, _ := http.NewRequest("GET", dependency.SSO_API_URL.String()+"/clients/OEE", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-Auth-SRPToken", result)
		q := request.URL.Query()
		q.Add("cluster", server.Cluster)
		q.Add("workspace", server.Workspace)
		q.Add("namespace", server.Namespace)
		q.Add("serviceName", "OEE")
		request.URL.RawQuery = q.Encode()
		response, _ := httpClient.Do(request)
		m, _ := simplejson.NewFromReader(response.Body)
		IFPToken = m.Get("clientSecret").MustString()
		fmt.Println("Cloud IFP Token:", IFPToken)
		time.Sleep(60 * time.Minute)
	}

}

func OnPremiseIFPToken() {

	for {
		httpClient := &http.Client{}
		content := map[string]string{"username": dependency.IFP_DESK_USERNAME, "password": dependency.IFP_DESK_PASSWORD}
		variable := map[string]interface{}{"input": content}
		httpRequestBody, _ := json.Marshal(map[string]interface{}{
			"query":     "mutation signIn($input: SignInInput!) {   signIn(input: $input) {     user {       name       __typename     }     __typename   } }",
			"variables": variable,
		})
		request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
		request.Header.Set("Content-Type", "application/json")
		response, _ := httpClient.Do(request)
		m, _ := simplejson.NewFromReader(response.Body)
		for {
			if len(m.Get("errors").MustArray()) == 0 {
				break
			} else {
				fmt.Println("retry refreshToken")
				httpRequestBody, _ = json.Marshal(map[string]interface{}{
					"query":     "mutation signIn($input: SignInInput!) { signIn(input: $input) { user { name __typename } __typename } }",
					"variables": variable,
				})
				request, _ = http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
				request.Header.Set("Content-Type", "application/json")
				response, _ = httpClient.Do(request)
				m, _ = simplejson.NewFromReader(response.Body)
				time.Sleep(6 * time.Minute)
			}
		}
		header := response.Header
		cookie := header["Set-Cookie"]
		var ifpToken, eiToken string
		for _, cookieContent := range cookie {
			tempSplit := strings.Split(cookieContent, ";")
			if strings.HasPrefix(tempSplit[0], "IFPToken") {
				ifpToken = tempSplit[0]
			} else if strings.HasPrefix(tempSplit[0], "EIToken") {
				eiToken = tempSplit[0]
			}
		}
		if eiToken == "" {
			IFPToken = ifpToken
		} else {
			IFPToken = ifpToken + ";" + eiToken
		}
		fmt.Println("On-Premise IFP Token:", IFPToken)
		time.Sleep(60 * time.Minute)
	}

}
