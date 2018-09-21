package main

import (
	"os"
	"encoding/csv"
	"sync"
)

// storeDataPerMonth loops through the map that contains the original csv file
// and separates the entries based on the month of the expiration date
func storeDataPerMonth(listMonths map[string][][]string) {
	var wg sync.WaitGroup
	var monthList = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	for _, month := range monthList {
		if listMonths[month] != nil {
			wg.Add(1)
			 go saveMonth(&wg , month, listMonths[month])
		}
	}
	wg.Wait()
}

// go routine for saving a month's data into its own csv file
func saveMonth(wg *sync.WaitGroup,month string, currentMonth [][]string) {
	defer wg.Done()
	var monthFileName = "./csvFiles/" + month + ".csv "
	file, err := os.Create(monthFileName)
	if err != nil {
		logger.Error("saveMonth Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	err = writer.WriteAll(currentMonth)
	if err != nil {
		logger.Error("saveMonth Cannot create file", err)
	}
	defer writer.Flush()
}
