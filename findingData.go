package main

import (
	"io"
	"os"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

type promotion struct {
	ID string
}

func getRecordById(id string) (result string) {
	result = "Not Found"
	files, err := ioutil.ReadDir("./csvFiles/")
	if err != nil {
		logger.Error("getRecordById: Directory", err)
	}

	for _, f := range files {
		if f.Size() != 0 {
			fmt.Println(f.Name())
			file, err := os.Open("./csvFiles/"+f.Name())
			if err != nil {
				logger.Errorf("readCsvFile Error:", err.Error())
			}
			// automatically call Close() at the end of current method
			defer file.Close()
			//
			reader := csv.NewReader(file)
			reader.Comma = ';'
			for {
				// read just one record
				record, err := reader.Read()
				// end-of-file is fitted into err
				if err == io.EOF {
					break
				} else if err != nil {
					logger.Error("getRecordById: File", err)
					return
				}
				if record[0] == id {
					fmt.Println("got it ")
					fmt.Println(record)
					result = strings.Join(record, " , ")
				}
			}
		}
	}
	return
}

// it should be used with local host
