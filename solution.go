package main

import (
	"github.com/excel-column-problem/router/v1"
	"net/http"
)

func main() {
	// Initialising routes for the http server and listening on port 9000
	router := apirouter.InitializeRouter()
	if err := http.ListenAndServe("localhost:9000", router); err != nil {
		panic(err.Error())
	}

}
