package main

import (
	"os"
	"encoding/csv"
	"io"
	"fmt"
)

type Record struct {
	Id              string
	Price           float64
	Expiration_date string
}

//
// ReadCSV will read our initial csv file and store it in an object
//
func ReadCSV() ( error) {
	// open file
	csvfile, err := os.Open("ids.csv")
	if err != nil {
		logger.Error(err)
		return err
	}
	defer csvfile.Close()

	//parse csv file
	reader := csv.NewReader(csvfile)
	cols := make(map[string]int)
	for rowCount := 0 ; ; rowCount++{
		record, err := reader.Read()
		if err == io.EOF {
			break
		}else if err != nil {
			logger.Error(err)
			return err
		}
		if rowCount==0{
			for index , col := range record{
				cols[col]= index
			}
		}else {
			rec, err := LoadedDataStoring(cols,record)
			if err != nil {
				logger.Errorf(err.Error())
			}
			fmt.Println(rec)
		}
	}
	return  nil
}
