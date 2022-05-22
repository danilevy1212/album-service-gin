package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumPostBody struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price"`
}

type AlbumPatchBody struct {
	Title  *string  `json:"title,omitempty"`
	Artist *string  `json:"artist,omitempty"`
	Price  *float64 `json:"price,omitempty"`
}
