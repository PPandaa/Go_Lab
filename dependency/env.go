package dependency

import (
	"GoLab/server"
	"GoLab/tool"
	"fmt"
	"net/url"
	"os"
)

var (
	SSO_API_URL               *url.URL
	IFP_DESK_API_URL          *url.URL
	IFP_DESK_CLIENT_SECRET    string
	IFP_DESK_USERNAME         string
	IFP_DESK_PASSWORD         string
	DAEMON_DATABROKER_API_URL *url.URL
)

func setENV() {

	logString := "Dependencies Info." + "\n"

	SSO_API_URL, _ = url.Parse(os.Getenv("SSO_API_URL"))
	logString += "  SSO_API_URL: " + SSO_API_URL.String() + "\n"

	ifps_desk_api_url := os.Getenv("IFP_DESK_API_URL")
	if tool.IsEmptyString(ifps_desk_api_url) {
		IFP_DESK_API_URL, _ = url.Parse("https://ifp-organizer-" + server.Namespace + "-" + server.Cluster + "." + server.External + "/graphql")
	} else {
		IFP_DESK_API_URL, _ = url.Parse(ifps_desk_api_url)
	}
	logString += "  IFP_DESK_API_URL: " + IFP_DESK_API_URL.String() + "\n"

	if server.Location == server.Cloud {
		IFP_DESK_CLIENT_SECRET = os.Getenv("IFP_DESK_CLIENT_SECRET")
		logString += "  IFP_DESK_CLIENT_SECRET: " + IFP_DESK_CLIENT_SECRET + "\n"
	} else {
		IFP_DESK_USERNAME = os.Getenv("IFP_DESK_USERNAME")
		logString += "  IFP_DESK_USERNAME: " + IFP_DESK_USERNAME + "\n"
		IFP_DESK_PASSWORD = os.Getenv("IFP_DESK_PASSWORD")
		logString += "  IFP_DESK_PASSWORD: " + IFP_DESK_PASSWORD + "\n"
	}

	daemon_databroker_api_url := os.Getenv(server.ServiceNameC + "_DAEMON_DATABROKER_API_URL")
	if tool.IsEmptyString(daemon_databroker_api_url) {
		DAEMON_DATABROKER_API_URL, _ = url.Parse("https://" + server.ServiceNameL + "-daemon-databroker-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	} else {
		DAEMON_DATABROKER_API_URL, _ = url.Parse(daemon_databroker_api_url)
	}
	logString += "  " + server.ServiceNameC + "_DAEMON_DATABROKER_API_URL: " + DAEMON_DATABROKER_API_URL.String() + "\n"

	fmt.Print(logString + "\n")

}
