package crud

import "carshop/internal/model"

var (
	Cars []model.Car
)

func init() {

	Cars = []model.Car{
		{
			Root: model.Root{
				ID:     "f43f0fe0-56c1-4f58-851e-9cfe351af748",
				Active: true,
			},
			Manufacturer: "Volkswagen",
			Model:        "Golf 8",
			Picture:      "https://www.b92.net/news/pics/2019/10/24/20399075025db17a1020862984778647_orig.jpg",
			Transmission: "automatic",
			Fuel:         "diesel",
			Type:         "hatchback",
			Price:        25000,
		},
		{
			Root: model.Root{
				ID:     "404dce22-aeeb-4aa7-b5d9-ba0eb0af4b75",
				Active: true,
			},
			Manufacturer: "Opel",
			Model:        "Insignia",
			Picture:      "https://www.b92.net/news/pics/2019/02/18/13438560765c6a6829c126f751480008_orig.jpg",
			Transmission: "manual",
			Fuel:         "diesel",
			Type:         "limousine",
			Price:        50000,
		},
		{
			Root: model.Root{
				ID:     "d44aa5cc-a4c5-4c56-a91d-0fc25263a6c6",
				Active: true,
			},
			Manufacturer: "Audi",
			Model:        "RS6",
			Picture:      "https://www.b92.net/news/pics/2019/08/21/13832399615d5cfa4175ce7418317815_orig.jpg",
			Transmission: "automatic",
			Fuel:         "petrol",
			Type:         "caravan",
			Price:        100000,
		},
		{
			Root: model.Root{
				ID:     "44d23ca6-59d8-4b9e-b7b1-3723f287d2c6",
				Active: true,
			},
			Manufacturer: "Dacia",
			Model:        "Duster",
			Picture:      "https://www.b92.net/news/pics/2021/06/28/124793220460d96fb7c1616566380933_1280x720x161600.jpg",
			Transmission: "manual",
			Fuel:         "gas",
			Type:         "SUV",
			Price:        30000,
		},
		{
			Root: model.Root{
				ID:     "ffac10ad-cb58-419a-9a5b-dec95bd908dc",
				Active: true,
			},
			Manufacturer: "Mercedes",
			Model:        "EQS",
			Picture:      "https://www.b92.net/news/pics/2021/09/06/12997866436135b8713cf3e318871538_1280x720x282700.jpg",
			Transmission: "automatic",
			Fuel:         "electric",
			Type:         "limousine",
			Price:        150000,
		},
		{
			Root: model.Root{
				ID:     "52dbe14d-fc14-4145-872f-00598ca64d86",
				Active: true,
			},
			Manufacturer: "Cupra",
			Model:        "Formentor",
			Picture:      "https://www.b92.net/news/pics/2020/10/07/11985151515f7db5ad4b95d911749173_orig.jpg",
			Transmission: "automatic",
			Fuel:         "petrol",
			Type:         "SUV",
			Price:        40000,
		},
		{
			Root: model.Root{
				ID:     "654c9f9e-8cbf-4b9c-9084-4c6381b8e854",
				Active: true,
			},
			Manufacturer: "Tesla",
			Model:        "Model S",
			Picture:      "https://www.b92.net/news/pics/2018/07/08/11902571865b41fc0a15d34554181640_orig.jpg",
			Transmission: "automatic",
			Fuel:         "electric",
			Type:         "limousine",
			Price:        100000,
		},
		{
			Root: model.Root{
				ID:     "fb1234e8-1f95-46b4-a727-099f5572a2e6",
				Active: true,
			},
			Manufacturer: "Renault",
			Model:        "Clio",
			Picture:      "https://www.b92.net/news/pics/2019/07/28/6538758695d3d456c8d1d3929243215_orig.jpg",
			Transmission: "manual",
			Fuel:         "diesel",
			Type:         "hatchback",
			Price:        20000,
		},
		{
			Root: model.Root{
				ID:     "834b125e-adc5-43f6-8fb4-e9ce24de8c80",
				Active: true,
			},
			Manufacturer: "Fiat",
			Model:        "500",
			Picture:      "https://www.b92.net/news/pics/2021/03/08/5379488716045e1f044165726017324_orig.jpg",
			Transmission: "manual",
			Fuel:         "gas",
			Type:         "hatchback",
			Price:        25000,
		},
		{
			Root: model.Root{
				ID:     "5986eb10-254b-493b-a5c6-ec3d06e77fd5",
				Active: true,
			},
			Manufacturer: "Mercedes",
			Model:        "E63",
			Picture:      "https://www.b92.net/news/pics/2017/02/02/155099174458930a2b176b0574653827_orig.jpg",
			Transmission: "automatic",
			Fuel:         "diesel",
			Type:         "caravan",
			Price:        75000,
		},
		{
			Root: model.Root{
				ID:     "a82ecf01-ee0c-4db4-b3a5-ba6f11646fb3",
				Active: true,
			},
			Manufacturer: "Alfa Romeo",
			Model:        "Giulia",
			Picture:      "https://www.b92.net/news/pics/2020/03/03/18365623605e5e0e20becb4549134738_orig.jpg",
			Transmission: "manual",
			Fuel:         "petrol",
			Type:         "limousine",
			Price:        75000,
		},
		{
			Root: model.Root{
				ID:     "02c7bcfd-dec6-4716-bd14-aae1416a0f80",
				Active: true,
			},
			Manufacturer: "Porsche",
			Model:        "Panamera",
			Picture:      "https://www.b92.net/news/pics/2020/08/26/14246548895f46704324aab244748797_orig.jpg",
			Transmission: "automatic",
			Fuel:         "petrol",
			Type:         "limousine",
			Price:        125000,
		},
		{
			Root: model.Root{
				ID:     "cf9f7233-97cf-4827-946b-2c10689a4314",
				Active: true,
			},
			Manufacturer: "Škoda",
			Model:        "Karoq",
			Picture:      "https://www.b92.net/news/pics/2017/12/30/4806489185a4756cfd47b2045437358_700x525x989800.jpg",
			Transmission: "manual",
			Fuel:         "petrol",
			Type:         "SUV",
			Price:        50000,
		},
		{
			Root: model.Root{
				ID:     "145fefcb-319f-44d7-b0d8-79b7ff2a8b6d",
				Active: true,
			},
			Manufacturer: "Volkswagen",
			Model:        "Arteon",
			Picture:      "https://www.b92.net/news/pics/2020/06/24/15671219255ef313b284685390372220_1280x720x12412300.jpg",
			Transmission: "automatic",
			Fuel:         "diesel",
			Type:         "limousine",
			Price:        50000,
		},
		{
			Root: model.Root{
				ID:     "e6a71212-14cf-467d-a2e4-903c94b7d685",
				Active: true,
			},
			Manufacturer: "Seat",
			Model:        "Ibiza",
			Picture:      "https://www.b92.net/news/pics/2021/04/15/18400060216078362962b2d328360520_orig.jpg",
			Transmission: "manual",
			Fuel:         "petrol",
			Type:         "hatchback",
			Price:        20000,
		},
	}
}
