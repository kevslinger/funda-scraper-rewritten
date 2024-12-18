package fundascraperrewritten_test

import (
	"fmt"
	"testing"

	fundascraperrewritten "github.com/kevslinger/funda-scraper-rewritten"
)

func TestGetRequestURL(t *testing.T) {
	area := "groningen"
	maxPrice := 350000
	minimumBedrooms := 3
	minimumSquareMeters := -1
	expected := fmt.Sprintf(`https://www.funda.nl/zoeken/koop/?selected_area=["%s"]&price="-%d"&bedrooms="%d-"`, area, maxPrice, minimumBedrooms)
	actual := fundascraperrewritten.GetRequestURL(area, maxPrice, minimumBedrooms, minimumSquareMeters)
	if expected != actual {
		t.Fatal(expected, actual)
	}
}
