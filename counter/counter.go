package counter

import (
	"fmt"
	"time"
)

func CountSheep(n int) {
	for n > 0 {
		fmt.Println("Sheep")
		time.Sleep(time.Second * 2)
		n--
	}
}
