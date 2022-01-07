package dependency

import (
	"fmt"
	"net/url"
	"os"

	"GoLab/server"
	"GoLab/tool"
)

var (
	ETCD_ENABLE               string
	ETCD_HOST                 string
	ETCD_PORT                 string
	ETCD_USERNAME             string
	ETCD_PASSWORD             string
	SSO_API_URL               *url.URL
	IFP_DESK_API_URL          *url.URL
	IFP_DESK_CLIENT_SECRET    string
	IFP_DESK_USERNAME         string
	IFP_DESK_PASSWORD         string
	UI_URL                    *url.URL
	API_URL                   *url.URL
	ETCD_BROKER_API_URL       *url.URL
	DAEMON_DATABROKER_API_URL *url.URL
)

func setENV() {

	logString := "Dependencies Info." + "\n"

	if tool.IsEmptyString(os.Getenv("ETCD_ENABLE")) {
		ETCD_ENABLE = "false"
	} else {
		ETCD_ENABLE = os.Getenv("ETCD_ENABLE")
	}
	logString += "  ETCD_ENABLE: " + ETCD_ENABLE + "\n"
	ETCD_HOST = os.Getenv("ETCD_HOST")
	logString += "  ETCD_HOST: " + ETCD_HOST + "\n"
	ETCD_PORT = os.Getenv("ETCD_PORT")
	logString += "  ETCD_PORT: " + ETCD_PORT + "\n"
	ETCD_USERNAME = os.Getenv("ETCD_USERNAME")
	logString += "  ETCD_USERNAME: " + ETCD_USERNAME + "\n"
	ETCD_PASSWORD = os.Getenv("ETCD_PASSWORD")
	logString += "  ETCD_PASSWORD: " + ETCD_PASSWORD + "\n"

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

	ui_url := os.Getenv(server.ServiceNameC + "_UI_URL")
	if tool.IsEmptyString(ui_url) {
		UI_URL, _ = url.Parse("https://" + server.ServiceNameL + "-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	} else {
		UI_URL, _ = url.Parse(ui_url)
	}
	logString += "  " + server.ServiceNameC + "_UI_URL: " + UI_URL.String() + "\n"

	api_url := os.Getenv(server.ServiceNameC + "_API_URL")
	if tool.IsEmptyString(api_url) {
		API_URL, _ = url.Parse("https://" + server.ServiceNameL + "-" + server.Namespace + "-api-" + server.Cluster + "." + server.External)
	} else {
		API_URL, _ = url.Parse(api_url)
	}
	logString += "  " + server.ServiceNameC + "_API_URL: " + API_URL.String() + "\n"

	etcd_broker_api_url := os.Getenv(server.ServiceNameC + "_ETCD_BROKER_API_URL")
	if len(etcd_broker_api_url) != 0 {
		ETCD_BROKER_API_URL, _ = url.Parse(etcd_broker_api_url)
	} else {
		ETCD_BROKER_API_URL, _ = url.Parse("https://" + server.ServiceNameL + "-etcd-broker-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	}
	logString += "  " + server.ServiceNameC + "_ETCD_BROKER_API_URL: " + ETCD_BROKER_API_URL.String() + "\n"

	daemon_databroker_api_url := os.Getenv(server.ServiceNameC + "_DAEMON_DATABROKER_API_URL")
	if tool.IsEmptyString(daemon_databroker_api_url) {
		DAEMON_DATABROKER_API_URL, _ = url.Parse("https://" + server.ServiceNameL + "-daemon-databroker-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	} else {
		DAEMON_DATABROKER_API_URL, _ = url.Parse(daemon_databroker_api_url)
	}
	logString += "  " + server.ServiceNameC + "_DAEMON_DATABROKER_API_URL: " + DAEMON_DATABROKER_API_URL.String() + "\n"

	fmt.Print(logString + "\n")

}
