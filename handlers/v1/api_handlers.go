package apihandlers

import (
	"encoding/json"
	"github.com/excel-column-problem/utility"
	"net/http"
	"strings"
)

func GetColumnTitleNo(w http.ResponseWriter, r *http.Request) {
	columnString, ok := r.URL.Query()["column-title"]
	// Checking for valid column title parameter is api request
	if !ok || len(columnString[0]) < 1 || !utility.IsStringAlphbatic(columnString[0]) {
		http.Error(w, "Bad request, kindly check parameters", http.StatusBadRequest)
		return
	}
	// To make sure column title string remain upper case
	columTitle := strings.ToUpper(columnString[0])
	columnNo := utility.ConvertColumnTitleToNo(columTitle)
	dataByte, _ := json.Marshal(columnNo)
	w.Write(dataByte)
	w.WriteHeader(http.StatusOK)
}
