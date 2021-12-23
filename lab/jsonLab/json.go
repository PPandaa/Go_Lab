package jsonLab

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type jsonStruct struct {
	Name     string `json:"name"`
	Birthday string `json:"bd"`
}

var stringInput = `{
	"name": "Peter", 
	"bd": "860502"
}`

var mapInput = map[string]interface{}{
	"name": "Peter",
	"bd":   "860502",
}

func StringToJSON() {

	newSimpleJSON, _ := simplejson.NewJson([]byte(stringInput))
	fmt.Println("SimpleJSON:", newSimpleJSON)

}

func StringToStruct() {

	var newJSON jsonStruct
	json.Unmarshal([]byte(stringInput), &newJSON)
	fmt.Printf("Struct: %+v\n", newJSON)

}

func MapToJSON() {

	temp, _ := json.Marshal(mapInput)
	newSimpleJSON, _ := simplejson.NewJson(temp)
	fmt.Println("SimpleJSON:", newSimpleJSON)

}
