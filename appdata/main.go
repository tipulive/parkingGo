package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	// Must print CGI header first
	fmt.Println("Content-Type: text/html")
	fmt.Println()

	fmt.Println("<html><body>")
	fmt.Println("<h1>Hello from Go CGI on OVH!</h1>")

	// Handle query parameters if any
	queryStr := os.Getenv("QUERY_STRING")
	if queryStr != "" {
		q, err := url.ParseQuery(queryStr)
		if err == nil {
			fmt.Println("<h2>Query Parameters:</h2><ul>")
			for k, v := range q {
				fmt.Printf("<li>%s = %v</li>\n", k, v)
			}
			fmt.Println("</ul>")
		} else {
			fmt.Println("<p>Error parsing query parameters</p>")
		}
	}

	fmt.Println("</body></html>")
}
