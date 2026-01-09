package dto

type ColorOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetColorsResponse struct {
	Colors []*ColorOutput `json:"colors"`
}
