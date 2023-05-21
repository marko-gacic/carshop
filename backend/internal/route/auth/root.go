package auth

import (
	"carshop/internal/model"
	"carshop/pkg/crypto"
)

var (
	Users  []model.User
	Tokens []model.Token
)

func init() {
	adminID, _ := crypto.UUID()
	guestID, _ := crypto.UUID()
	userID, _ := crypto.UUID()

	Users = []model.User{
		{
			Root: model.Root{
				ID: adminID,
			},
			Role:        "admin",
			Username:    "admin@carshop.com",
			Password:    "admin",
			Salutation:  "Mrs",
			FirstName:   "Lisa",
			LastName:    "Jones",
			Married:     true,
			PhoneNumber: "+38164123123",
		},
		{
			Root: model.Root{
				ID: userID,
			},
			Role:        "user",
			Username:    "user@carshop.com",
			Password:    "user",
			Salutation:  "Mr",
			FirstName:   "John",
			LastName:    "Thompson",
			Married:     false,
			PhoneNumber: "+38164321321",
			Address: model.Address{
				Country: "Serbia",
				Street:  "Mise Dimitrijevica 1",
				ZipCode: 21000,
				City:    "Novi Sad",
			},
		},
		{
			Root: model.Root{
				ID: guestID,
			},
			Role:        "guest",
			Username:    "guest@carshop.com",
			Password:    "guest",
			Salutation:  "Mr",
			FirstName:   "Tom",
			LastName:    "Pearson",
			Married:     false,
			PhoneNumber: "+38164231231",
			Address: model.Address{
				Country: "Serbia",
				Street:  "Knez Mihailova",
				ZipCode: 11000,
				City:    "Beograd",
			},
		},
	}

	Tokens = make([]model.Token, 0)
}
