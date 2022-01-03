package main

import (
	"github.com/gofiber/fiber/v2"
)

//Obviously doesn't need the struct in the keyword
type BatterStruct struct {
	Batter []map[string]string `json:"batter"`
	Toppings []map[string]string `json:"toppings"`
}

type DataStruct struct {
	Id string `json: "dataId"`
	Batters []BatterStruct `json:"batters"`
}

func main(){
	app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {

		data := DataStruct{
			Id: "0001",
			Batters: []BatterStruct {
				{
					Batter: []map[string]string {
						{
							"id": "1001", 
							"type": "Regular",
						},
						{
							"id": "1002", 
							"type": "Chocolate",
						},
					},
					Toppings: []map[string]string {
						{ "id": "5001", "type": "None" },
						{ "id": "5002", "type": "Glazed" },
						{ "id": "5005", "type": "Sugar" },
						{ "id": "5007", "type": "Powdered Sugar" },
						{ "id": "5006", "type": "Chocolate with Sprinkles" },
						{ "id": "5003", "type": "Chocolate" },
						{ "id": "5004", "type": "Maple" },
					},
				},
			},
		}

		return c.JSON(&fiber.Map{
			"response": data,
		})

  })

  app.Listen(":3000")
}
