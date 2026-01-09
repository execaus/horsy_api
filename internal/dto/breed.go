package dto

type BreedOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetBreedsResponse struct {
	Breeds []*BreedOutput `json:"breeds"`
}
