package uuidLab

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func Run() {

	// V1 基於時間
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u1.String())

	// V4 基於亂數
	u4 := uuid.New()
	fmt.Println(u4.String()) // a0d99f20-1dd1-459b-b516-dfeca4005203

}
