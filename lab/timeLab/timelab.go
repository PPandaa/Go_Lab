package timeLab

import (
	"fmt"
	"time"
)

func Run() {

	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
	fmt.Println(formatted)

}
