package mongodb

import (
	"GoLab/guard"
	"GoLab/pkg"
	"GoLab/server"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"gopkg.in/mgo.v2"
)

var (
	DB          *mgo.Database
	Session     *mgo.Session
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

	logString := "Mongo Database Info." + "\n"

	if server.Location == server.Cloud {
		ensaasService := os.Getenv("ENSAAS_SERVICES")
		if !pkg.IsEmptyString(ensaasService) {
			tempReader := strings.NewReader(ensaasService)
			m, _ := simplejson.NewFromReader(tempReader)
			mongodb := m.Get("mongodb").GetIndex(0).Get("credentials").MustMap()
			MongodbInfo.URL = mongodb["externalHosts"].(string)
			MongodbInfo.Database = mongodb["database"].(string)
			MongodbInfo.Username = mongodb["username"].(string)
			MongodbInfo.Password = mongodb["password"].(string)
		} else {
			MongodbInfo.URL = os.Getenv("MONGODB_URL")
			MongodbInfo.Database = os.Getenv("MONGODB_DATABASE")
			MongodbInfo.Username = os.Getenv("MONGODB_USERNAME")
			MongodbInfo.Password = os.Getenv("MONGODB_PASSWORD")
		}
	} else {
		MongodbInfo.URL = os.Getenv("MONGODB_URL")
		MongodbInfo.Database = os.Getenv("MONGODB_DATABASE")
		MongodbInfo.Username = os.Getenv("MONGODB_USERNAME")
		MongodbInfo.AuthDatabase = os.Getenv("MONGODB_AUTH_SOURCE")
		mongodbPasswordFile := os.Getenv("MONGODB_PASSWORD_FILE")
		if !pkg.IsEmptyString(mongodbPasswordFile) {
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
	logString += "  URL: " + MongodbInfo.URL + "\n" +
		"  Database: " + MongodbInfo.Database + "\n" +
		"  Username: " + MongodbInfo.Username + "\n" +
		"  Password: " + MongodbInfo.Password + "\n"
	fmt.Print(logString + "\n")

}

func Connect() {

	newSession, err := mgo.Dial(MongodbInfo.URL)
	if err != nil {
		guard.Logger.Error("Database Connect Fail -> " + err.Error())
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
			guard.Logger.Fatal("Database Login Fail -> " + err.Error())
		}
	} else {
		DB = Session.DB(MongodbInfo.AuthDatabase)
		err = DB.Login(MongodbInfo.Username, MongodbInfo.Password)
		if err != nil {
			guard.Logger.Fatal("Database Login Fail -> " + err.Error())
		}
		DB = Session.DB(MongodbInfo.Database)
	}
	guard.Logger.Info("Database Connect Success")

}

func ConnectCheck() {

	err := Session.Ping()
	if err != nil {
		guard.Logger.Error("Database Connect Check Fail")
		Session.Refresh()
		guard.Logger.Info("Database Reconnect \n")
	}

}
