package dependency

import (
	"GoLab/server"
	"GoLab/tool"

	"fmt"
	"net/url"
	"os"
	"strconv"
)

var (
	SSO_API_URL *url.URL

	AppNameC   string
	AppNameL   string
	AppVersion string

	UI_URL          *url.URL
	Is_UI_Reachable bool

	API_URL          *url.URL
	Is_API_Reachable bool

	ETCD_BROKER_API_URL          *url.URL
	Is_ETCD_BROKER_API_Reachable bool

	DAEMON_DATABROKER_API_URL          *url.URL
	Is_DAEMON_DATABROKER_API_Reachable bool

	IFP_DESK_UI_URL           *url.URL
	Is_IFP_DESK_UI_Reachable  bool
	IFP_DESK_API_URL          *url.URL
	Is_IFP_DESK_API_Reachable bool
	IFP_DESK_CLIENT_SECRET    string
	IFP_DESK_USERNAME         string
	IFP_DESK_PASSWORD         string

	ServiceSecret string
)

func Set() {

	from_env()

}

func from_env() {

	logString := "Dependencies Info." + "\n"

	SSO_API_URL, _ = url.Parse(os.Getenv("SSO_API_URL"))
	logString += "  SSO_API_URL: " + SSO_API_URL.String() + "\n"

	AppVersion = os.Getenv("IAPP_VERSION")
	logString += "  IAPP_VERSION: " + AppVersion + "\n"

	AppNameC = os.Getenv("IAPP_NAME_CAPITAL")
	logString += "  IAPP_NAME_CAPITAL: " + AppNameC + "\n"

	AppNameL = os.Getenv("IAPP_NAME_LOWER")
	logString += "  IAPP_NAME_LOWER: " + AppNameL + "\n"

	ui_url := os.Getenv(server.AppNameC + "_UI_URL")
	if tool.IsEmptyString(ui_url) {
		UI_URL, _ = url.Parse("https://" + server.AppNameL + "-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	} else {
		UI_URL, _ = url.Parse(ui_url)
	}
	Is_UI_Reachable = tool.IsSiteReachable(UI_URL.String())
	logString += "  UI_URL: " + UI_URL.String() + " ..... " + strconv.FormatBool(Is_UI_Reachable) + "\n"

	api_url := os.Getenv(server.AppNameC + "_API_URL")
	if tool.IsEmptyString(api_url) {
		API_URL, _ = url.Parse("https://" + server.AppNameL + "-" + server.Namespace + "-api-" + server.Cluster + "." + server.External)
	} else {
		API_URL, _ = url.Parse(api_url)
	}
	Is_API_Reachable = tool.IsSiteReachable(API_URL.String())
	logString += "  API_URL: " + API_URL.String() + " ..... " + strconv.FormatBool(Is_API_Reachable) + "\n"

	etcd_broker_api_url := os.Getenv(server.AppNameC + "_ETCD_BROKER_API_URL")
	if len(etcd_broker_api_url) != 0 {
		ETCD_BROKER_API_URL, _ = url.Parse(etcd_broker_api_url)
	} else {
		ETCD_BROKER_API_URL, _ = url.Parse("https://" + server.AppNameL + "-etcd-broker-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	}
	Is_ETCD_BROKER_API_Reachable = tool.IsSiteReachable(ETCD_BROKER_API_URL.String())
	logString += "  " + server.AppNameC + "_ETCD_BROKER_API_URL: " + ETCD_BROKER_API_URL.String() + " ..... " + strconv.FormatBool(Is_ETCD_BROKER_API_Reachable) + "\n"

	daemon_databroker_api_url := os.Getenv(server.AppNameC + "_DAEMON_DATABROKER_API_URL")
	if tool.IsEmptyString(daemon_databroker_api_url) {
		DAEMON_DATABROKER_API_URL, _ = url.Parse("https://" + server.AppNameL + "-daemon-databroker-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	} else {
		DAEMON_DATABROKER_API_URL, _ = url.Parse(daemon_databroker_api_url)
	}
	Is_DAEMON_DATABROKER_API_Reachable = tool.IsSiteReachable(DAEMON_DATABROKER_API_URL.String())
	logString += "  " + server.AppNameC + "_DAEMON_DATABROKER_API_URL: " + DAEMON_DATABROKER_API_URL.String() + " ..... " + strconv.FormatBool(Is_DAEMON_DATABROKER_API_Reachable) + "\n"

	ifps_desk_ui_url := os.Getenv("IFP_DESK_UI_URL")
	if tool.IsEmptyString(ifps_desk_ui_url) {
		IFP_DESK_UI_URL, _ = url.Parse("https://ifp-organizer-" + server.Namespace + "-" + server.Cluster + "." + server.External)
	} else {
		IFP_DESK_UI_URL, _ = url.Parse(ifps_desk_ui_url)
	}
	Is_IFP_DESK_UI_Reachable = tool.IsSiteReachable(IFP_DESK_UI_URL.String())
	logString += "  IFP_DESK_UI_URL: " + IFP_DESK_UI_URL.String() + " ..... " + strconv.FormatBool(Is_IFP_DESK_UI_Reachable) + "\n"

	ifps_desk_api_url := os.Getenv("IFP_DESK_API_URL")
	if tool.IsEmptyString(ifps_desk_api_url) {
		IFP_DESK_API_URL, _ = url.Parse("https://ifp-organizer-" + server.Namespace + "-" + server.Cluster + "." + server.External + "/graphql")
	} else {
		IFP_DESK_API_URL, _ = url.Parse(ifps_desk_api_url)
	}
	Is_IFP_DESK_API_Reachable = tool.IsSiteReachable(IFP_DESK_API_URL.String())
	logString += "  IFP_DESK_API_URL: " + IFP_DESK_API_URL.String() + " ..... " + strconv.FormatBool(Is_IFP_DESK_API_Reachable) + "\n"

	if server.Location == server.Cloud {
		IFP_DESK_CLIENT_SECRET = os.Getenv("IFP_DESK_CLIENT_SECRET")
		logString += "  IFP_DESK_CLIENT_SECRET: " + IFP_DESK_CLIENT_SECRET + "\n"
	} else {
		IFP_DESK_USERNAME = os.Getenv("IFP_DESK_USERNAME")
		logString += "  IFP_DESK_USERNAME: " + IFP_DESK_USERNAME + "\n"
		IFP_DESK_PASSWORD = os.Getenv("IFP_DESK_PASSWORD")
		logString += "  IFP_DESK_PASSWORD: " + IFP_DESK_PASSWORD + "\n"
	}

	fmt.Print(logString + "\n")

}
