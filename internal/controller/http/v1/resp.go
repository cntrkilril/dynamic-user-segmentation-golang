package v1

import "github.com/gofiber/fiber/v2"

func newResp(data any, errorArray []error) fiber.Map {
	messages := make([]string, 0, len(errorArray))
	for _, v := range errorArray {
		messages = append(messages, v.Error())
	}
	return fiber.Map{
		"data":   data,
		"errors": messages,
	}
}
