package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://crd.buyinsta.com/me?id=1")
	if err != nil {
		fmt.Printf("Error parsing the url: %v", err)
	}

	fmt.Println(u.String(), u.RawQuery)

	// add or delete or update new values to query params manually
	newQueryParams := u.Query() // keeps existing ones
	// newQueryParams := url.Values{} // creates new Query param values
	newQueryParams.Add("name", "raju")
	u.RawQuery = newQueryParams.Encode()

	fmt.Println(u.String(), u.RawQuery)
}
