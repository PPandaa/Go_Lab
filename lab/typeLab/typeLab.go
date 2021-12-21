package typeLab

import (
	"encoding/json"
	"fmt"
)

func Run() {

	temp := 1
	fmt.Print(float64(temp))

}

type User struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func FromInterfaceToStruct() {

	me := map[string]interface{}{"name": "Peter", "birthday": "86/05/02"}
	var result User
	temp, _ := json.Marshal(me)
	json.Unmarshal(temp, &result)
	fmt.Println(result)

}
