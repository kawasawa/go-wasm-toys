package entity

type GenerateRandomRequest struct {
	Characters string
	Length     int
	Count      int
}

type GenerateRandomResponse struct {
	Randoms []string
}
