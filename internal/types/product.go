package types

type Person struct {
	Name string `json:"name" yaml:"name"`
}

type Category struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type CategoryLadder struct {
	Ladder []Category `json:"ladder" yaml:"ladder"`
	Root   string     `json:"root" yaml:"root"`
}

type Product struct {
	Title            string `json:"title" yaml:"title"`
	Subtitle         string `json:"subtitle" yaml:"subtitle"`
	ReleaseDate      string `json:"release_date" yaml:"release_date"`
	PublisherName    string `json:"publisher_name" yaml:"publisher_name"`
	PublisherSummary string `json:"publisher_summary" yaml:"publisher_summary"`

	Authors   []Person `json:"authors" yaml:"authors"`
	Narrators []Person `json:"narrators" yaml:"narrators"`

	ProductImages struct {
		Image500 string `json:"500" yaml:"500"`
	} `json:"product_images" yaml:"product_images"`

	CategoryLadders []CategoryLadder `json:"category_ladders" yaml:"category_ladders"`
}

type AudibleResponse struct {
	Product Product `json:"product"`
}
