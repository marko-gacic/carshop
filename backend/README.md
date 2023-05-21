**CAR SHOP**

Prerequisite:
_install GoLang locally_

This is a simple CRUD backend service with mocked data, that wraps simple operations including:

2. Create operation (add car)
3. Read operation (retrieve 1 car from given ID)
4. List operation (retrieve X cars)
7. Delete operation (soft-delete 1 row from database, from given ID)

Also, this service offers authentication services:

1. Log the user in (create token)
2. Log the user out (destroy token)

## MODELS ##

type Car struct {
	Root
	Manufacturer string       `json:"manufacturer"`
	Model        string       `json:"model"`
	Picture      string       `json:"picture"`
	Transmission string       `json:"transmission"`
	Fuel         string       `json:"fuel"`
	Type         string       `json:"type"`
	Price        float64      `json:"price"`
}

type Root struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`
}

type User struct {
	Root
	Role        string  `json:"role"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Salutation  string  `json:"salutation"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Married     bool    `json:"married"`
	PhoneNumber string  `json:"phoneNumber"`
	Address     Address `json:"address"`
}

type Address struct {
	Country string `json:"country"`
	Street  string `json:"street"`
	ZipCode int64  `json:"zipCode"`
	City    string `json:"city"`
}
