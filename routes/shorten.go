package routes

import (
	"time"
	"github.com/gofiber/fiber/v2"
)

type request struct{
	URL 			string 				`json:"url"`
	CustomShort 	string 				`json:"short"`
	Expiry 			time.Duration 		`json:"expiry"`
}

type response struct{
	URL 			string 			`json:"url"`
	CustomShort 	string 			`json:"short"`
	Expiry 			time.Duration 	`json:"expiry"`
	XRateRemaining  int				`json:"rate_limit"`
	XRateLimitReset time.Duration	`json:"rate_limit_reset"`
} 

func shortenURL(c *fiber.Ctx) error{
	body := new(request)

	if err := c.BodyParser(&body);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse Json"})
	}
	

}
