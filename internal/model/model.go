package model

type Profile struct {
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Name string `json:"category_name"`
	Hubungan []Hubung `json:"hubungan"`
}

type Hubung struct {
	Name string `json:"name"`
	Link string `json:"link"`
}