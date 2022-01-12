package miniodb

import (
	"GoLab/guard"
	"GoLab/tool"

	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/minio/minio-go"
)

var (
	IsMinioEnable = false
	Client        *minio.Client
	valueFrom     string
	MinioDBInfo   InfoStruct
	Bucket        = "peter-test"
)

type InfoStruct struct {
	IsSSL     bool
	URL       string
	AccessKey string
	SecretKey string
}

func Set() {

	logString := "MinIO Info." + "\n"

	valueFrom = "ENV"
	if !tool.IsEmptyString(os.Getenv("STORAGE_API_URL")) {
		IsMinioEnable = true

		url := strings.Split(os.Getenv("STORAGE_API_URL"), "://")
		urlHead := url[0]
		if urlHead == "https" {
			MinioDBInfo.IsSSL = true
		} else {
			MinioDBInfo.IsSSL = false
		}

		if !tool.IsEmptyString(os.Getenv("STORAGE_API_PORT")) {
			MinioDBInfo.URL = url[1] + ":" + os.Getenv("STORAGE_API_PORT")
		} else {
			MinioDBInfo.URL = url[1]
		}

		MinioDBInfo.AccessKey = os.Getenv("STORAGE_ACCESS_KEY")
		MinioDBInfo.SecretKey = os.Getenv("STORAGE_SECRET_KEY")

		logString += "  FROM: " + valueFrom + "\n" +
			"    SSL: " + strconv.FormatBool(MinioDBInfo.IsSSL) + "\n" +
			"    URL: " + MinioDBInfo.URL + "\n" +
			"    AccessKey: " + MinioDBInfo.AccessKey + "\n" +
			"    SecretKey: " + MinioDBInfo.SecretKey + "\n"
	}

	fmt.Print(logString + "\n")

}

func Connect() {

	var err error

	if IsMinioEnable {
		Client, err = minio.New(MinioDBInfo.URL, MinioDBInfo.AccessKey, MinioDBInfo.SecretKey, MinioDBInfo.IsSSL)
		if err != nil {
			guard.Logger.Fatal("minio connect fail -> " + err.Error())
		} else {
			guard.Logger.Info("minio connect success")
			Client.MakeBucket(Bucket, "")
		}
	}

}

// func creatBucket() {

// 	bucket := "peter-test"
// 	err := Client.MakeBucket(bucket, "")
// 	if err != nil {
// 		exists, errBucketExists := Client.BucketExists(bucket)
// 		if errBucketExists == nil && exists {
// 			guard.Logger.Info(bucket + " Bucket Already Exist")
// 		} else {
// 			guard.Logger.Fatal(err.Error())
// 		}
// 	} else {
// 		guard.Logger.Info(bucket + " Bucket Created")
// 	}

// }
