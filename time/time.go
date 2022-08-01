package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC3339Nano, "2022-10-22T10:20:04Z")
	// t, err := time.Parse("2006-01-02", "2022-10-22") // can also be given manually but need to default given date and time in go to parse
	if err != nil {
		fmt.Printf("Error parsing the string to time: %v", err)
	}
	fmt.Println(t.String())

	// now time can be formatted as required but needs to use the default date and time to change the format
	fmt.Println(t.Format(time.RubyDate))
	fmt.Println(t.Format("02-01-2006"))
}
