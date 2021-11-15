package desk

import (
	"GoLab/auth"
	"GoLab/dependency"
	"GoLab/guard"
	"GoLab/server"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func RegisterOutbound() {

	httpClient := &http.Client{}
	content := map[string]interface{}{"name": server.ServiceNameC, "sourceId": "scada_" + server.ServiceNameL, "url": dependency.DAEMON_DATABROKER_API_URL.String(), "active": true}
	variable := map[string]interface{}{"input": content}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query":     "mutation ($input: AddOutboundInput!) { addOutbound(input: $input) { outbound { id name url sourceId allowUnauthorized active connected } } }",
		"variables": variable,
	})
	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}
	request.Header.Set("Content-Type", "application/json")
	response, _ := httpClient.Do(request)
	if response.StatusCode == 200 {
		dependency.IsDeskEnable = true
	} else {
		dependency.IsDeskEnable = false
	}
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) == 0 {
		guard.Logger.Info("Register Outbound " + server.ServiceNameC + " Success")
	} else {
		guard.Logger.Info("Outbound " + server.ServiceNameC + " is already exist")
	}

}
