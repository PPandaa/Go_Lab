package mongodbLab

// import (
// 	"GoLab/database/mongodb"
// 	"fmt"

// 	"gopkg.in/mgo.v2/bson"
// )

// const (
// 	TestCollection = "iii.peter.test"
// )

// func Insert() {

// 	data := map[string]interface{}{
// 		"Faculty": "Computer Science",
// 		"ID":      "00457029",
// 		"Name":    "Peter",
// 	}
// 	collection := mongodb.DB.C(TestCollection)
// 	err := collection.Insert(data)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(data)
// 	}

// }

// func Pipe() {

// 	collection := mongodb.DB.C(TestCollection)
// 	result := map[string]interface{}{}
// 	filter := []bson.M{
// 		{"$match": bson.M{"Name": "Shift"}},
// 	}
// 	err := collection.Pipe(filter).One(&result)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// }

// func Update() {

// 	collection := mongodb.DB.C(TestCollection)
// 	uData := map[string]interface{}{
// 		"Name": "H5T2H",
// 	}
// 	filter := bson.M{
// 		"ID": "00457029",
// 	}
// 	change := bson.M{"$set": uData}
// 	err := collection.Update(filter, change)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(filter, uData)
// 	}

// }

// func Remove() {

// 	collection := mongodb.DB.C(TestCollection)
// 	filter := bson.M{"ID": "00457029"}
// 	changeInfo, err := collection.RemoveAll(filter)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(changeInfo)
// 	}

// }

// func RemoveAllCollection() {

// 	collectionNames, _ := mongodb.DB.CollectionNames()
// 	for _, collectionName := range collectionNames {
// 		fmt.Println("Remove Collection ->", collectionName)
// 		collection := mongodb.DB.C(collectionName)
// 		collection.DropCollection()
// 	}

// }

// func FindDistinctValue() {

// 	collection := mongodb.DB.C(TestCollection)
// 	var result []string
// 	collection.Find(nil).Distinct("ID", &result)
// 	fmt.Println(result)

// }
