package main

import (
	"fmt"
	service "github.com/ykim135/DateApi/service"
	"time"
)

func main() {
	var allFirstMonday []time.Time
	var month = time.Month(1)

	output := service.GetAllFirstMonday(allFirstMonday, month)

	fmt.Println("============ First Mondays ============")
	for _, element := range output {
		fmt.Printf("%d-%02d-%02d\n",
			element.Year(), element.Month(), element.Day())
	}
}
