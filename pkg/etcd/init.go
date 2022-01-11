package etcd

import (
	"GoLab/dependency"
	"GoLab/guard"
	"GoLab/server"
	"GoLab/tool"

	"net/http"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

func Run() {

	etcd_enable, etcd_host, etcd_port := get_etcd_info()

	if etcd_enable {
		etcdConn := &EtcdConn{}
		etcdConn.new(etcd_enable, etcd_host, etcd_port, "", "")

		etcdCliForIApp := &EtcdCli{}
		etcdCliForIApp.client = etcdCliForIApp.connect(etcdConn)
		etcdCliForIApp.start("iapps/", dependency.AppNameC, dependency.AppVersion)
		var iAppLink IAppLink
		iAppLink.new(dependency.AppNameC, dependency.UI_URL.String(), dependency.API_URL.String()+"/images/icon.svg", "iframe")
		etcdCliForIApp.putIAppLink(iAppLink)
		var iAppFeature IAppFeature
		iAppFeature.new(dependency.AppNameC, dependency.API_URL.String()+"/images/icon.svg")
		etcdCliForIApp.putIAppFeature(iAppFeature)
		etcdCliForIApp.startElection()

		etcdCliForService := &EtcdCli{}
		etcdCliForService.client = etcdCliForService.connect(etcdConn)
		etcdCliForService.start("services/", dependency.AppNameC, dependency.AppVersion)
		etcdCliForService.startElection()

		go etcdCliForService.watchServiceSecrets()
	} else {
		guard.Logger.Fatal("ETCD disable")
	}

}

func get_etcd_info() (bool, string, string) {

	etcd_enable := false
	etcd_host := ""
	etcd_port := ""

	for {
		if dependency.Is_IFP_DESK_API_Reachable {
			request, _ := http.NewRequest("GET", dependency.IFP_DESK_UI_URL.String()+"/simple-metadata/etcd3Hosts", nil)
			response, _ := server.HttpClient.Do(request)
			if response.StatusCode != 404 {
				m, _ := simplejson.NewFromReader(response.Body)
				etcd_url := m.GetIndex(0).MustString() // http://rtm-ifp-etcd:2379
				temp := strings.Split(etcd_url, ":")
				etcd_enable = true
				etcd_host = temp[0] + ":" + temp[1]
				etcd_port = temp[2]
			}
			break
		} else {
			guard.Logger.Error("Desk unreachable")
			dependency.Is_IFP_DESK_API_Reachable = tool.IsSiteReachable(dependency.IFP_DESK_UI_URL.String())
		}
		time.Sleep(1 * time.Minute)
	}

	return etcd_enable, etcd_host, etcd_port

}
