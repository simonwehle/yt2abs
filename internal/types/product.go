package types

type Person struct {
	Name string `json:"name"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryLadder struct {
	Ladder []Category `json:"ladder"`
	Root   string     `json:"root"`
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

	CategoryLadders []CategoryLadder `json:"category_ladders"`
}

type AudibleResponse struct {
	Product Product `json:"product"`
}