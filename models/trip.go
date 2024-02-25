package models

type Trip struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Score           string `json:"score"`
	Hour            string `json:"hour"`
	Local           string `json:"local"`
	Price           string `json:"price"`
	DepartureCity   string `json:"departureCity"`
	DestinationCity string `json:"destinationCity"`
	Day             string `json:"day"`
}
