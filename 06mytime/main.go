package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time pkg of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdTime := time.Date(2023, time.November, 29, 22, 11, 0, 0, time.UTC)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("02-01-2006 Monday"))
}
