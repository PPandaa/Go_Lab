package tool

func CheckType(i interface{}) string {

	switch i.(type) {
	case int:
		return "int"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "none"
	}

}

func InterfaceListToStringList(iList []interface{}) []string {

	stringList := []string{}

	for _, v := range iList {
		stringList = append(stringList, v.(string))
	}

	return stringList

}
