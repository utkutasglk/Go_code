package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userInfo@]host[:port[/path][?query][#fragment]

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		fmt.Println("Error parsing URL :", err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port :", parsedUrl.Port())
	fmt.Println("Raw Query :", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	rawURL := "https://deneme.com/path?name=Utku&age=30"

	parsedUrl1, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL :", err)
		return
	}

	queryParams := parsedUrl1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	// building URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "utku.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	query.Set("name", "Utku")
	baseUrl.RawQuery = query.Encode()

	fmt.Println("Built URL :", baseUrl.String())

	values := url.Values{}

	// Add key-value pairs to the values object
	values.Add("name","Jane")
	values.Add("age","30")
	values.Add("city","Istanbul")
	values.Add("country","TR")

	// Encode
	encodedQuery := values.Encode()

	fmt.Println(encodedQuery)

	// Build a URL
	baseURL1 := "https://example.com/search"
	fullUrl := baseURL1 + "?" + encodedQuery

	fmt.Println(fullUrl)

}
