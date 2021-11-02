package mongodbLab

import (
	"GoLab/database/mongodb"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func PipeTest() {

	collection := mongodb.DB.C("iii.mom.All")
	result := map[string]interface{}{}
	filter := []bson.M{
		{"$match": bson.M{"Name": "Shift"}},
	}
	err := collection.Pipe(filter).One(&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
