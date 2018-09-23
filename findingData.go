package main

import (
	"io"
	"os"
	"encoding/csv"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"sync"
	"strings"
)

type channelFiles struct {
	Values [][]string
	Id     string
}

// GetRecordById takes ID from parameters and call function readDirectoryFiles to find it
// calls the template for display
func GetRecordById(w http.ResponseWriter, r *http.Request) {
	defer elapsed()()
	vars := mux.Vars(r)
	id := vars["id"]
	result := readDirectoryFiles(id)
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "index.html", result)
	if err != nil {
		logger.Error("couldn't execute template", err)
	}
}

// readDirectoryFiles load file in a directory
// and call a go routine to read each file
// giving an ID in the parameter it returns either the record for this ID
// or not found
func readDirectoryFiles(idToLookFor string) (result string) {
	var asyncWaitGroup = sync.WaitGroup{}
	var monthValues = make(map[string][][]string)
	result = "Not Found"
	files, err := ioutil.ReadDir("./csvFiles/")
	if err != nil {
		logger.Error("getRecordById: Directory", err)
	}
	asyncWaitGroup.Add(len(files))
	filesChan := make(chan channelFiles, len(files))
	go func() {
		asyncWaitGroup.Wait()
		close(filesChan)
	}()
	for _, file := range files {
		go loopThroughFile(file, filesChan)
	}
	for channelItem := range filesChan {
		monthValues[channelItem.Id] = channelItem.Values
		asyncWaitGroup.Done()
	}
	var resultList = make(map[string]string)
	for month, monthValue := range monthValues {
		resultList[month] = result
		for _, record := range monthValue {
			if record[0] == idToLookFor {
				result = strings.Join(record, " , ")
			}
		}
	}
	return
}

// loopThroughFile goroutine reads csv  file and adds it to channel
func loopThroughFile(file os.FileInfo, files chan channelFiles) {
	var filesChan channelFiles
	var fileValues [][]string
	if file.Size() != 0 {
		file, err := os.Open("./csvFiles/" + file.Name())
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
			fileValues = append(fileValues, record)
		}
	}
	filesChan.Values = fileValues
	filesChan.Id = file.Name()
	files <- filesChan

}
