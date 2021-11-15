package dependency

import (
	"GoLab/server"
	"fmt"
	"net/url"
	"os"
)

var (
	SSO_API_URL               *url.URL
	IFP_DESK_API_URL          *url.URL
	DAEMON_DATABROKER_API_URL *url.URL
)

func setEnv() {

	logString := "Dependencies Info." + "\n"

	SSO_API_URL, _ = url.Parse(os.Getenv("SSO_API_URL"))
	logString += "  SSO_API_URL: " + SSO_API_URL.String() + "\n"

	ifps_desk_api_url := os.Getenv("IFP_DESK_API_URL")
	if len(ifps_desk_api_url) != 0 {
		IFP_DESK_API_URL, _ = url.Parse(ifps_desk_api_url)
	} else {
		IFP_DESK_API_URL, _ = url.Parse("https://ifp-organizer-" + server.Namespace + "-" + server.Cluster + "." + server.External + "/graphql")
	}
	logString += "  IFP_DESK_API_URL: " + IFP_DESK_API_URL.String() + "\n"

	daemon_databroker_api_url := os.Getenv(server.ServiceNameC + "_DAEMON_DATABROKER_API_URL")
	if len(daemon_databroker_api_url) != 0 {
		DAEMON_DATABROKER_API_URL, _ = url.Parse(daemon_databroker_api_url)
	} else {
		DAEMON_DATABROKER_API_URL, _ = url.Parse("https://" + server.ServiceNameL + "-daemon-databroker-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	}
	logString += "  " + server.ServiceNameC + "_DAEMON_DATABROKER_API_URL: " + DAEMON_DATABROKER_API_URL.String() + "\n"

	fmt.Print(logString + "\n")

}
