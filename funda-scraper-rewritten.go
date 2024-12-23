package fundascraperrewritten

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const defaultIntFlagValue = -1

func GetRequestURL(area string, maximumPrice int, minimumBedrooms int, minimumSquareMeters int) string {
	var searchUrl strings.Builder
	searchUrl.WriteString(fmt.Sprintf(`https://www.funda.nl/zoeken/koop/?selected_area=["%s"]`, area))
	if maximumPrice > defaultIntFlagValue {
		searchUrl.WriteString(fmt.Sprintf(`&price="-%d"`, maximumPrice))
	}
	if minimumBedrooms > defaultIntFlagValue {
		searchUrl.WriteString(fmt.Sprintf(`&bedrooms="%d-"`, minimumBedrooms))
	}
	if minimumSquareMeters > defaultIntFlagValue {
		searchUrl.WriteString(fmt.Sprintf(`&floor_area="%d-"`, minimumSquareMeters))
	}
	return searchUrl.String()
}

func GetListingsFromResponse(response *http.Response) ([]string, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return make([]string, 0), err
	}
	responseBody := string(body)
	fundaHttpRegex := regexp.MustCompile(`https://www.funda.nl/detail/koop`)
	listingIndices := fundaHttpRegex.FindAllStringIndex(responseBody, -1)
	uniqueListings := make(map[string]bool)
	for _, listingIndex := range listingIndices {
		endIdx := listingIndex[1]
		for responseBody[endIdx] != '"' {
			endIdx++
		}
		uniqueListings[responseBody[listingIndex[0]:endIdx]] = true
	}
	listingsSlice := make([]string, 0)
	for listing := range uniqueListings {
		listingsSlice = append(listingsSlice, listing)
	}
	return listingsSlice, nil
}

func Main() int {
	searchArea := flag.String("area", "nl", "Name of the area to look for houses. Eg: nl, amsterdam, utrecht (nb: must be lowercase)")
	maximumPrice := flag.Int("max-price", defaultIntFlagValue, "Maximum price. Eg: 500000, 375000")
	minimumBedrooms := flag.Int("min-bedrooms", defaultIntFlagValue, "Minimum number of bedrooms. Eg: 2")
	minimumSquareMeters := flag.Int("min-square-meters", defaultIntFlagValue, "Minmium size (in square meters). Eg: 75")
	flag.Parse()

	requestUrl := GetRequestURL(*searchArea, *maximumPrice, *minimumBedrooms, *minimumSquareMeters)
	request, err := prepareRequest(requestUrl)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return 1
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error performing HTTP request: %v\n", err)
		return 1
	}
	listings, err := GetListingsFromResponse(response)
	if err != nil {
		fmt.Printf("Error getting listings from HTTP response: %v\n", err)
		return 1
	}
	fmt.Println(listings)
	return 0
}

func prepareRequest(requestUrl string) (*http.Request, error) {
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
	return req, nil
}
