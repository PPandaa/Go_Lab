package server

import (
	"GoLab/tool"

	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

const (
	AppNameC    = "IFPS_III"
	AppNameL    = "ifps-iii"
	ServiceName = "lab"
	Cloud       = "Cloud"
	OnPremise   = "On-Premise"
)

var (
	Location              string
	Datacenter            string
	Workspace             string
	Cluster               string
	Namespace             string
	External              string
	IsEnsaasServiceEnable = false
	EnsaasService         *simplejson.Json
	LastWaconnTime        time.Time
	HttpClient            = &http.Client{}
)

func Check() {

	logString := "Server Info." + "\n"

	Datacenter = os.Getenv("datacenter")
	Workspace = os.Getenv("workspace")
	Cluster = os.Getenv("cluster")
	Namespace = os.Getenv("namespace")
	External = os.Getenv("external")

	if !tool.IsEmptyString(Datacenter) {
		Location = Cloud
		logString += "  Location: " + Location + "\n" +
			"  Datacenter: " + Datacenter + "\n" +
			"  Workspace: " + Workspace + "\n" +
			"  Cluster: " + Cluster + "\n" +
			"  Namespace: " + Namespace + "\n" +
			"  External: " + External + "\n"

		ensaasService := os.Getenv("ENSAAS_SERVICES")
		if !tool.IsEmptyString(ensaasService) {
			tempReader := strings.NewReader(ensaasService)
			EnsaasService, _ = simplejson.NewFromReader(tempReader)
			IsEnsaasServiceEnable = true
		}
	} else {
		Location = OnPremise
		logString += "  Location: " + Location + "\n"
	}

	fmt.Print(logString + "\n")

}
