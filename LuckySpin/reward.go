package LuckySpin

import "github.com/gofiber/fiber"

type Fruit struct {
	Name     string
	Id       int
	Quantity int
}

var nonFruit = Fruit{
	Name: "Quả nịt",
	Id: 9999,
	Quantity: 9999999,
}

var fruits = []Fruit{
	{Name: "táo", Id: 1001, Quantity: 1},
	{Name: "chanh", Id: 1002, Quantity: 1},
	{Name: "dưa hấu", Id: 1003, Quantity: 2},
	{Name: "dứa", Id: 1004, Quantity: 1},
}

type Response struct {
	Message string      `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
	Status  int         `json:"Status,omitempty"`
}

func WriteSuccess(c *fiber.Ctx, v interface{}) error {
	res := Response{
		Message: "Success",
		Data:    v,
		Status:  200,
	}

	// Return
	return c.JSON(res)
}