package intermediate

import (
	"fmt"
	"net/url"
)

func url_parsing() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing the url:", err)
		return
	}

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw query:", parsedUrl.RawQuery)
	fmt.Println("Fragement:", parsedUrl.Fragment)

	rawUrl1 := "https://example.com/path?name=john&age=30"

	parseUrl , err1 := url.Parse(rawUrl1)
	if err1 != nil {
		fmt.Println("Error parsing the url:", err1)
		return

	}
	queryParams := parseUrl.Query()
	fmt.Println(queryParams)
	fmt.Println("Name :", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))
}
