package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

var logger = logrus.Logger{}

func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Println("Excution time : %v\n", time.Since(start))
	}
}

func main() {
	defer elapsed()()
	router := mux.NewRouter().StrictSlash(false)
	router.Methods("GET").Path("/get/{id}").HandlerFunc(LoadData)
	http.Handle("/", router)
	fmt.Println("listening on ", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
