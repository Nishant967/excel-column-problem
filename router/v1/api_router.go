package apirouter

import (
	"github.com/excel-column-problem/handlers/v1"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/get-column-no",apihandlers.GetColumnTitleNo).Methods("GET")
	return router
}
