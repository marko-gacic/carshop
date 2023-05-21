package model

type Role string

const (
	Empty     Role = ""
	UserGuest Role = "guest"
	UserAdmin Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	Root
	Role        Role    `json:"role"`
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
