package usecase

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}
