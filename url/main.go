package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://www.example.com:8080/path/to/resource?kye1=value&key2=value"
	fmt.Printf("Type of url is: %T\n", myUrl)

	parsedURL, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Printf("Type of URL is: %T\n", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Port:", parsedURL.Port())

	// Modifying the URL
	parsedURL.Path = "/new/path"
	parsedURL.RawQuery = "newkey=newValue"
	fmt.Println("Modified URL:", parsedURL.String())

	// Adding a new query parameter
	queryParams := parsedURL.Query() // parsedURL.Query() will Parse the existing query.
	queryParams.Add("newKey", "newValue")

	parsedURL.RawQuery = queryParams.Encode() // Encode the query parameters back to a string

	fmt.Println("URL with new query parameter:", parsedURL.String())
}
