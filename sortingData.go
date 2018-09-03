package main

import (
	"sync"
	"fmt"
	"time"
	"net/http"
	"strconv"
)

var waitGroup sync.WaitGroup
var data chan string

type chanelRes struct {
	Values *Record
	Id     string
	error  error
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	//var structResult chanelRes
	fmt.Println("waiting ")
	time.Sleep(10 * time.Millisecond)
	ReadCSV()
}

func LoadedDataStoring(cols map[string]int, record []string) (*Record, error) {
	var r = Record{}
	col := cols["id"]
	id  ,err := strconv.Atoi(record[col])
	if err != nil {
		return nil , err
	}
	fmt.Println("kfp ", id )
	return  &r , nil
}
