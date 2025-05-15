package audible

import (
	"encoding/json"
	"fmt"
	"net/http"

	"yt2abs/internal/types"
)

func FetchMetadata(asin string) (*types.Product, error) {
	url  := fmt.Sprintf("https://api.audible.com/1.0/catalog/products/%s?response_groups=media,product_extended_attrs,category_ladders", asin)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; audible-scraper/1.0)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("faulty server response: %d", resp.StatusCode)
	}

	var data types.AudibleResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data.Product, nil
}