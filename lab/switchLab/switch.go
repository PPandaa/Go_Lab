package switchLab

import "fmt"

func Run() {

	gender := "male"
	switch gender {
	case "male":
		fmt.Println("男")
	case "female":
		fmt.Println("女")
	default:
		fmt.Println("")
	}

}
