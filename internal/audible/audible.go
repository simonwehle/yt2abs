package audible

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
}

type Product struct {
	Title            string `json:"title"`
	Subtitle         string `json:"subtitle"`
	ReleaseDate      string `json:"release_date"`
	PublisherName    string `json:"publisher_name"`
	PublisherSummary string `json:"publisher_summary"`

	Authors          []Person `json:"authors"`
	Narrators        []Person `json:"narrators"`

	ProductImages struct {
		Image500 string `json:"500"`
	} `json:"product_images"`
}

type AudibleResponse struct {
	Product Product `json:"product"`
}

func FetchMetadata(asin string) (*Product, error) {
	url  := fmt.Sprintf("https://api.audible.com/1.0/catalog/products/%s?response_groups=media,product_extended_attrs", asin)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set fake User-Agent to avoid blocking
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; audible-scraper/1.0)")
	// Optionally set Accept if needed
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("faulty server response: %d", resp.StatusCode)
	}

	var data AudibleResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data.Product, nil
}