package etcd

import (
	"GoLab/dependency"
)

func Initial() {

	uiUrl := dependency.UI_URL.String() + "/index"
	iconUrl := dependency.DAEMON_DATABROKER_API_URL.String() + "/images/icon.svg"

	etcdConn := &EtcdConn{}
	etcdConn.new(dependency.ETCD_ENABLE, dependency.ETCD_HOST, dependency.ETCD_PORT, dependency.ETCD_USERNAME, dependency.ETCD_PASSWORD)

	etcdCliForIapp := &EtcdCli{}
	etcdCliForIapp.start(etcdConn, true, "GoLab", "1.0.0")

	var iAppLink IAppLink
	iAppLink.new("GoLab", uiUrl, iconUrl, "iframe")
	etcdCliForIapp.putIAppLink(iAppLink)

	var iAppFeature IAppFeature
	iAppFeature.new("GoLab", iconUrl)
	etcdCliForIapp.putIAppFeature(iAppFeature)

	etcdCliForIapp.startElection()

}
