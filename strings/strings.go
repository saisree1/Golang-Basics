package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// converting a different data type to string and vice versa
	intData := 1

	intDataAsString := strconv.Itoa(intData)
	// strconv.ParseFloat() // different data types conversion exists
	fmt.Println(intDataAsString == "1")

	intDataTemp, err := strconv.Atoi(intDataAsString)
	// intDataTemp, err := strconv.Atoi("Raju") // gives error
	if err != nil {
		fmt.Println("Error converting from string to int : ", err)
	}
	fmt.Println(intDataTemp == 1)

	// string functions
	stringData := "Ramu"

	fmt.Println(strings.ContainsAny(stringData, "am"))
	fmt.Println(strings.Contains(stringData, "am"))
	fmt.Println(strings.ToLower(stringData), strings.ToUpper(stringData))
	fmt.Println(strings.Trim(stringData, "u"))
	fmt.Println(strings.EqualFold(stringData, "ramu"), strings.EqualFold(stringData, "RaMU"))

	// string builder faster and doesnt require copies to add strings
	var stringBuilderData strings.Builder
	stringBuilderData.WriteString(stringData)
	stringBuilderData.WriteString(stringData)
	fmt.Println(stringBuilderData.String())
}
