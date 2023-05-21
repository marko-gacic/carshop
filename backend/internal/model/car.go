package model

type Car struct {
	Root
	Manufacturer string       `json:"manufacturer"`
	Model        string       `json:"model"`
	Picture      string       `json:"picture"`
	Transmission Transmission `json:"transmission"`
	Fuel         Fuel         `json:"fuel"`
	Type         Type         `json:"type"`
	Price        float64      `json:"price"`
}

type Transmission string

const (
	Manual    Transmission = "manual"
	Automatic Transmission = "automatic"
)

type Fuel string

const (
	Diesel   Fuel = "diesel"
	Petrol   Fuel = "petrol"
	Gas      Fuel = "gas"
	Electric Fuel = "electric"
)

type Type string

const (
	Hatchback Type = "hatchback"
	Limousine Type = "limousine"
	Caravan   Type = "caravan"
	SUV       Type = "SUV"
)
