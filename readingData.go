package main

import (
	"os"
	"encoding/csv"
	"io"
	"fmt"
	"time"
	"net/http"
)

// ReadCSV will read our initial csv file and store it in an object
func GetInitialData(w http.ResponseWriter, r *http.Request) {
	defer elapsed()()
	var listMonths = make(map[string][][]string)
	file, err := os.Open("ids.csv")
	if err != nil {
		logger.Errorf("readCsvFile Error:", err.Error())
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	lineCount := 0
	for {
		// read just one record
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("GetInitialData Error:", err)
			return
		}
		// record is an array of string so is directly printable
		listMonths = getMonth(listMonths, record)
		lineCount += 1
	}
	storeDataPerMonth(listMonths)
	logger.Info("done loading")
}

// getMonth out of the expiration date and add the current
// input into its corresponding map according to the month
func getMonth(listMonths map[string][][]string, record []string) (map[string][][]string) {
	t, err := time.Parse("2006-01-02", record[2][0:10])
	if err != nil {
		logger.Error("parsing time ", err)
	}
	month := t.Month().String()
	var array = listMonths[month]
	array = append(array, record)
	listMonths[month] = array
	return listMonths
}
