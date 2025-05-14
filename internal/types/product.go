package types

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