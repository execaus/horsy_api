package dto

type GeneticMarkerOutput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetGeneticMarkersResponse struct {
	GeneticMarkers []*GeneticMarkerOutput `json:"geneticMarkers"`
}
