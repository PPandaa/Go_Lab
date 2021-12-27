package mongodb

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"GoLab/guard"
	"GoLab/server"
	"GoLab/tool"

	"gopkg.in/mgo.v2"
)

var (
	DB          *mgo.Database
	Session     *mgo.Session
	valueFrom   string
	MongodbInfo InfoStruct
)

type InfoStruct struct {
	URL          string
	Database     string
	Username     string
	Password     string
	AuthDatabase string
}

func Set() {

	logString := "MongoDB Info." + "\n"

	if server.Location == server.Cloud {
		if server.IsEnsaasServiceEnable && len(server.EnsaasService.Get("mongodb").MustArray()) != 0 {
			valueFrom = "ENSAAS_SERVICE"
			mongodb := server.EnsaasService.Get("mongodb").GetIndex(0).Get("credentials").MustMap()
			MongodbInfo.URL = mongodb["externalHosts"].(string)
			MongodbInfo.Database = mongodb["database"].(string)
			MongodbInfo.Username = mongodb["username"].(string)
			MongodbInfo.Password = mongodb["password"].(string)
		} else {
			valueFrom = "ENV"
			MongodbInfo.URL = os.Getenv("MONGODB_URL")
			MongodbInfo.Database = os.Getenv("MONGODB_DATABASE")
			MongodbInfo.Username = os.Getenv("MONGODB_USERNAME")
			MongodbInfo.Password = os.Getenv("MONGODB_PASSWORD")
		}
	} else {
		valueFrom = "ENV"
		MongodbInfo.URL = os.Getenv("MONGODB_URL")
		MongodbInfo.Database = os.Getenv("MONGODB_DATABASE")
		MongodbInfo.Username = os.Getenv("MONGODB_USERNAME")
		MongodbInfo.AuthDatabase = os.Getenv("MONGODB_AUTH_SOURCE")
		mongodbPasswordFile := os.Getenv("MONGODB_PASSWORD_FILE")
		if !tool.IsEmptyString(mongodbPasswordFile) {
			mongodbPassword, err := ioutil.ReadFile(mongodbPasswordFile)
			if err != nil {
				guard.Logger.Sugar().Fatalw("MongoDB Password File", "FilePath", mongodbPasswordFile)
			} else {
				MongodbInfo.Password = string(mongodbPassword)
			}
		} else {
			MongodbInfo.Password = os.Getenv("MONGODB_PASSWORD")
		}
	}
	logString += "  FROM: " + valueFrom + "\n" +
		"    URL: " + MongodbInfo.URL + "\n" +
		"    Database: " + MongodbInfo.Database + "\n" +
		"    Username: " + MongodbInfo.Username + "\n" +
		"    Password: " + MongodbInfo.Password + "\n"

	fmt.Print(logString + "\n")

}

func Connect() {

	newSession, err := mgo.Dial(MongodbInfo.URL)
	if err != nil {
		guard.Logger.Error("MongoDB Connect Fail -> " + err.Error() + "\n")
		for err != nil {
			newSession, err = mgo.Dial(MongodbInfo.URL)
			time.Sleep(5 * time.Second)
		}
	}
	Session = newSession

	if server.Location == server.Cloud {
		DB = Session.DB(MongodbInfo.Database)
		err = DB.Login(MongodbInfo.Username, MongodbInfo.Password)
		if err != nil {
			guard.Logger.Fatal("MongoDB Connect Fail -> " + err.Error())
		}
	} else {
		DB = Session.DB(MongodbInfo.AuthDatabase)
		err = DB.Login(MongodbInfo.Username, MongodbInfo.Password)
		if err != nil {
			guard.Logger.Fatal("MongoDB Connect Fail -> " + err.Error())
		}
		DB = Session.DB(MongodbInfo.Database)
	}

	guard.Logger.Info("MongoDB Connect Success")

}

func ConnectCheck() {

	err := Session.Ping()
	if err != nil {
		guard.Logger.Error("MongoDB Connect Check Fail")
		Session.Refresh()
		guard.Logger.Info("MongoDB Reconnect")
	}

}
