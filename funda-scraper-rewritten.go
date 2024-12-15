package fundascraperrewritten

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const (
	fundaSearchBuyUrl          = `https://www.funda.nl/zoeken/koop/?selected_area=["nl"]`
	fundaBuyHouseResultPattern = `https://www.funda.nl/detail/koop`
	userAgentHeader            = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"
)

func Main() int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fundaSearchBuyUrl, nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return 1
	}
	req.Header.Add("user-agent", userAgentHeader)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error performing HTTP request: %v\n", err)
		return 1
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return 1
	}
	bodyStr := string(body)
	fundaHttpRegex := regexp.MustCompile(fundaBuyHouseResultPattern)
	listingIndices := fundaHttpRegex.FindAllStringIndex(bodyStr, -1)
	listings := make([]string, 0)
	for _, listingIndex := range listingIndices {
		endIdx := listingIndex[1]
		for bodyStr[endIdx] != '"' {
			endIdx++
		}
		listings = append(listings, bodyStr[listingIndex[0]:endIdx])
	}
	fmt.Println(listings)
	return 0
}
