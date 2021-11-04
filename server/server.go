package server

import (
	"GoLab/pkg"
	"fmt"
	"os"
)

const (
	Cloud     = "Cloud"
	OnPremise = "On-Premise"
)

var (
	Location string

	Datacenter string
	Workspace  string
	Cluster    string
	Namespace  string
	External   string
)

func Set() {

	logString := "Server Info." + "\n"

	Datacenter = os.Getenv("datacenter")
	Workspace = os.Getenv("workspace")
	Cluster = os.Getenv("cluster")
	Namespace = os.Getenv("namespace")
	External = os.Getenv("external")

	if !pkg.IsEmptyString(Datacenter) {
		Location = Cloud
		logString += "  Location: " + Location + "\n" +
			"  Datacenter: " + Datacenter + "\n" +
			"  Workspace: " + Workspace + "\n" +
			"  Cluster: " + Cluster + "\n" +
			"  Namespace: " + Namespace + "\n" +
			"  External: " + External + "\n"
	} else {
		Location = OnPremise
		logString += "  Location: " + Location + "\n"
	}

	fmt.Print(logString + "\n")

}
