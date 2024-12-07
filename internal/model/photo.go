package model

type Photo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	ImageData   []byte `json:"image_data"`
}
