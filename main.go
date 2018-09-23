package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

var logger = logrus.Logger{}
var tpl *template.Template

func init() {
	var err error
	tpl = template.Must(template.ParseGlob("./views/*.html"))
	tpl, err = tpl.ParseFiles("./views/index.html")
	if err != nil {
		logger.Error("Init : couldn't parse files in template", err)
	}
}
func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Excution time : %v\n", time.Since(start))
	}
}

func main() {
	//go func() {
	//	c := time.Tick(30 * time.Minute)
	//	for range c {
	//		GetInitialData()
	//	}
	//}()
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Methods("GET").Path("/loadFile").Name("Load").HandlerFunc(GetInitialData)
	r.Methods("GET").Path("/promotions/{id}").Name("Find").HandlerFunc(GetRecordById)
	http.ListenAndServe(":8080", nil)
}
