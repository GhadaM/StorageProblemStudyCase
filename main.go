package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"fmt"
)

var logger = logrus.Logger{}

func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Excution time : %v\n", time.Since(start))
	}
}

func main() {
	defer elapsed()()
	//GetInitialData()
	res := getRecordById("172FFC14-D229-4C93-B06B-F48B8C095512")
	fmt.Println(res)
}
