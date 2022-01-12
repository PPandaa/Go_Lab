package dependency

import (
	"GoLab/database/miniodb"
	"GoLab/database/mongodb"
	"GoLab/database/redisdb"
	"GoLab/server"
	"GoLab/tool"

	"fmt"
	"net/url"
	"os"
	"strconv"
)

var (
	SSO_API_URL *url.URL

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

	fmt.Print("Dependencies Info." + "\n")

	from_env()

	mongodb.Set()
	mongodb.Connect()

	redisdb.Set()
	redisdb.Connect()

	miniodb.Set()
	miniodb.Connect()

}

func from_env() {

	logString := ""

	SSO_API_URL, _ = url.Parse(os.Getenv("SSO_API_URL"))
	logString += "  SSO_API_URL: " + SSO_API_URL.String() + "\n"

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
