package server

import (
	"GoLab/tool"

	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bitly/go-simplejson"
)

const (
	ServiceName = "lab"
	Cloud       = "Cloud"
	OnPremise   = "On-Premise"

	DefaultAppNameC = "IFPS_III"
	DefaultAppNameL = "ifps-iii"
	DefaultVersion  = "1.0.0"
)

var (
	HttpClient = &http.Client{}
	Location   string

	Datacenter string
	Workspace  string
	Cluster    string
	Namespace  string
	External   string

	EnsaasService         *simplejson.Json
	IsEnsaasServiceEnable bool

	AppNameC   string
	AppNameL   string
	AppVersion string

	UI_URL                    *url.URL
	API_URL                   *url.URL
	AUTH_API_URL              *url.URL
	COMMON_API_URL            *url.URL
	DATASOURCE_API_URL        *url.URL
	ETCD_BROKER_API_URL       *url.URL
	DAEMON_DATABROKER_API_URL *url.URL
)

func Up() {

	fmt.Print("Server Info." + "\n")
	check_server_location()
	check_service()

}

func check_server_location() {

	logString := ""

	Datacenter = os.Getenv("datacenter")
	Workspace = os.Getenv("workspace")
	Cluster = os.Getenv("cluster")
	Namespace = os.Getenv("namespace")
	External = os.Getenv("external")

	if !tool.IsEmptyString(Datacenter) {
		Location = Cloud
		logString += "  Location: " + Location + "\n" +
			"    Datacenter: " + Datacenter + "\n" +
			"    Workspace: " + Workspace + "\n" +
			"    Cluster: " + Cluster + "\n" +
			"    Namespace: " + Namespace + "\n" +
			"    External: " + External + "\n"

		ensaasService := os.Getenv("ENSAAS_SERVICES")
		if !tool.IsEmptyString(ensaasService) {
			IsEnsaasServiceEnable = true
			tempReader := strings.NewReader(ensaasService)
			EnsaasService, _ = simplejson.NewFromReader(tempReader)
		} else {
			IsEnsaasServiceEnable = false
		}
	} else {
		Location = OnPremise
		logString += "  Location: " + Location + "\n"
	}

	fmt.Print(logString + "\n")

}

func check_service() {

	logString := ""

	// iapp info
	appNameC := os.Getenv("IAPP_NAME_CAPITAL")
	if tool.IsEmptyString(appNameC) {
		AppNameC = DefaultAppNameC
	} else {
		AppNameC = appNameC
	}
	logString += "  IAPP_NAME_CAPITAL: " + AppNameC + "\n"

	appNameL := os.Getenv("IAPP_NAME_LOWER")
	if tool.IsEmptyString(appNameL) {
		AppNameL = DefaultAppNameL
	} else {
		AppNameL = appNameL
	}
	logString += "  IAPP_NAME_LOWER: " + AppNameL + "\n"

	appVersion := os.Getenv("IAPP_VERSION")
	if tool.IsEmptyString(appVersion) {
		AppVersion = DefaultVersion
	} else {
		AppVersion = appVersion
	}
	logString += "  IAPP_VERSION: " + AppVersion + "\n" + "\n"

	// iapp service info
	ui_url := os.Getenv(AppNameC + "_UI_URL")
	if tool.IsEmptyString(ui_url) {
		UI_URL, _ = url.Parse("https://" + AppNameL + "-" + Namespace + "-" + Cluster + "." + External)
	} else {
		UI_URL, _ = url.Parse(ui_url)
	}
	logString += "  UI_URL: " + UI_URL.String() + "\n"

	api_url := os.Getenv(AppNameC + "_API_URL")
	if tool.IsEmptyString(api_url) {
		API_URL, _ = url.Parse("https://" + AppNameL + "-api-" + Namespace + "-" + Cluster + "." + External)
	} else {
		API_URL, _ = url.Parse(api_url)
	}
	logString += "  API_URL: " + API_URL.String() + "\n"

	auth_api_url := os.Getenv(AppNameC + "_AUTH_API_URL")
	if tool.IsEmptyString(auth_api_url) {
		AUTH_API_URL, _ = url.Parse("https://" + AppNameL + "-auth-" + Namespace + "-" + Cluster + "." + External)
	} else {
		AUTH_API_URL, _ = url.Parse(auth_api_url)
	}
	logString += "  AUTH_API_URL: " + AUTH_API_URL.String() + "\n"

	common_api_url := os.Getenv(AppNameC + "_COMMON_API_URL")
	if tool.IsEmptyString(common_api_url) {
		COMMON_API_URL, _ = url.Parse("https://" + AppNameL + "-common-api-" + Namespace + "-" + Cluster + "." + External)
	} else {
		COMMON_API_URL, _ = url.Parse(common_api_url)
	}
	logString += "  COMMON_API_URL: " + COMMON_API_URL.String() + "\n"

	datasource_api_url := os.Getenv(AppNameC + "_DATASOURCE_API_URL")
	if tool.IsEmptyString(datasource_api_url) {
		DATASOURCE_API_URL, _ = url.Parse("https://" + AppNameL + "-datasource-" + Namespace + "-" + Cluster + "." + External)
	} else {
		DATASOURCE_API_URL, _ = url.Parse(datasource_api_url)
	}
	logString += "  DATASOURCE_API_URL: " + DATASOURCE_API_URL.String() + "\n"

	etcd_broker_api_url := os.Getenv(AppNameC + "_ETCD_BROKER_API_URL")
	if len(etcd_broker_api_url) != 0 {
		ETCD_BROKER_API_URL, _ = url.Parse(etcd_broker_api_url)
	} else {
		ETCD_BROKER_API_URL, _ = url.Parse("https://" + AppNameL + "-etcd-broker-" + Namespace + "-" + Cluster + "." + External)
	}
	logString += "  ETCD_BROKER_API_URL: " + ETCD_BROKER_API_URL.String() + "\n"

	daemon_databroker_api_url := os.Getenv(AppNameC + "_DAEMON_DATABROKER_API_URL")
	if tool.IsEmptyString(daemon_databroker_api_url) {
		DAEMON_DATABROKER_API_URL, _ = url.Parse("https://" + AppNameL + "-daemon-databroker-" + Namespace + "-" + Cluster + "." + External)
	} else {
		DAEMON_DATABROKER_API_URL, _ = url.Parse(daemon_databroker_api_url)
	}
	logString += "  DAEMON_DATABROKER_API_URL: " + DAEMON_DATABROKER_API_URL.String() + "\n"

	fmt.Print(logString + "\n")

}
