package utility

import (
	"regexp"
)

// Regex to make sure that only alphabets are there in column title parameter
var IsStringAlphbatic = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// Function to convert column title string to number
func ConvertColumnTitleToNo(columnTitle string) int {
	columnNo := 0
	for i := 0; i < len(columnTitle); i++ {
		// Used base as 64 considering alphabet will be upper case
		columnNo = int(columnTitle[i]) - 64 + (columnNo * 26)
	}
	return columnNo
}
