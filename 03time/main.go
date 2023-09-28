package main

import (
	"fmt"
	"time"
	// "time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
}
