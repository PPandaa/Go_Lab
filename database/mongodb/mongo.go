// mgo.v2 是一個 已經停止維護 (deprecated) 的 MongoDB Go 驅動，它不支援 MongoDB 4.2 以上的版本，包括 MongoDB 8.0。
// 因此，你不能用 mgo.v2 連接 mongo:8.0，因為它不支援 MongoDB 的新協議和身份驗證方式。

package mongodb

// import (
// 	"GoLab/guard"
// 	"GoLab/server"
// 	"GoLab/tool"

// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"time"

// 	"gopkg.in/mgo.v2"
// )

// var (
// 	DB          *mgo.Database
// 	MongodbInfo InfoStruct
// 	Session     *mgo.Session
// 	valueFrom   string
// )

// type InfoStruct struct {
// 	URL          string
// 	Database     string
// 	Username     string
// 	Password     string
// 	AuthDatabase string
// }

// func Set() {

// 	logString := "  MongoDB Info." + "\n"

// 	if server.Location == server.Cloud {
// 		if server.IsEnsaasServiceEnable && len(server.EnsaasService.Get("mongodb").MustArray()) != 0 {
// 			valueFrom = "ENSAAS_SERVICE"
// 			mongodb := server.EnsaasService.Get("mongodb").GetIndex(0).Get("credentials").MustMap()
// 			MongodbInfo.URL = mongodb["externalHosts"].(string)
// 			MongodbInfo.Database = mongodb["database"].(string)
// 			MongodbInfo.Username = mongodb["username"].(string)
// 			MongodbInfo.Password = mongodb["password"].(string)
// 		} else {
// 			valueFrom = "ENV"
// 			MongodbInfo.URL = os.Getenv("MONGODB_URL")
// 			MongodbInfo.Database = os.Getenv("MONGODB_DATABASE")
// 			MongodbInfo.Username = os.Getenv("MONGODB_USERNAME")
// 			MongodbInfo.Password = os.Getenv("MONGODB_PASSWORD")
// 		}
// 	} else {
// 		valueFrom = "ENV"
// 		MongodbInfo.URL = os.Getenv("MONGODB_URL")
// 		MongodbInfo.Database = os.Getenv("MONGODB_DATABASE")
// 		MongodbInfo.Username = os.Getenv("MONGODB_USERNAME")
// 		MongodbInfo.AuthDatabase = os.Getenv("MONGODB_AUTH_SOURCE")
// 		mongodbPasswordFile := os.Getenv("MONGODB_PASSWORD_FILE")
// 		if !tool.IsEmptyString(mongodbPasswordFile) {
// 			mongodbPassword, err := ioutil.ReadFile(mongodbPasswordFile)
// 			if err != nil {
// 				guard.Logger.Sugar().Fatalw("mongodb password file", "file path", mongodbPasswordFile)
// 			} else {
// 				MongodbInfo.Password = string(mongodbPassword)
// 			}
// 		} else {
// 			MongodbInfo.Password = os.Getenv("MONGODB_PASSWORD")
// 		}
// 	}
// 	logString += "    FROM: " + valueFrom + "\n" +
// 		"      URL: " + MongodbInfo.URL + "\n" +
// 		"      Database: " + MongodbInfo.Database + "\n" +
// 		"      Username: " + MongodbInfo.Username + "\n" +
// 		"      Password: " + MongodbInfo.Password + "\n"

// 	fmt.Print(logString + "\n")

// }

// func Connect() {

// 	newSession, err := mgo.Dial(MongodbInfo.URL)
// 	if err != nil {
// 		guard.Logger.Fatal("mongodb connect fail -> " + err.Error() + "\n")
// 		for err != nil {
// 			newSession, err = mgo.Dial(MongodbInfo.URL)
// 			time.Sleep(5 * time.Second)
// 		}
// 	}
// 	Session = newSession

// 	if server.Location == server.Cloud {
// 		DB = Session.DB(MongodbInfo.Database)
// 		err = DB.Login(MongodbInfo.Username, MongodbInfo.Password)
// 		if err != nil {
// 			guard.Logger.Fatal("mongodb connect fail -> " + err.Error())
// 		}
// 	} else {
// 		DB = Session.DB(MongodbInfo.AuthDatabase)
// 		err = DB.Login(MongodbInfo.Username, MongodbInfo.Password)
// 		if err != nil {
// 			guard.Logger.Fatal("mongodb connect fail -> " + err.Error())
// 		}
// 		DB = Session.DB(MongodbInfo.Database)
// 	}

// }

// func ConnectCheck() {

// 	err := Session.Ping()
// 	if err != nil {
// 		guard.Logger.Error("mongodb connect check fail")
// 		Session.Refresh()
// 		guard.Logger.Info("mongodb reconnect")
// 	} else {
// 		guard.Logger.Info("mongodb connect check success")
// 	}

// }
