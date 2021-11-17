package stringLab

import (
	"encoding/json"
	"fmt"
)

func Show() {

	strMap := make(map[string]interface{})
	str := `{"appName":"ifp.org","modelName":"Group","id":"R3JvdXA.YZSUhUAekgAGtQbM","name":"AAA","description":"","timeZone":"Pacific/Rarotonga","createdAt":"2021-11-17T05:35:01.491Z","updatedAt":"2021-11-17T05:35:01.491Z"}`
	json.Unmarshal([]byte(str), &strMap)
	fmt.Println(strMap)

}
