package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/luisfer00/url-shortener/pkg/services"
)

type insertURLResponseBody struct {
	Url string `json:"url" validate:"url"`
}

func validateInsertURL(body *insertURLResponseBody) []string {
	var errorsSlice []string

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorsSlice = append(errorsSlice, err.Error())
		}
	}

	return errorsSlice
}

func GetURLController(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Param is empty",
		})
	}

	url, err := services.GetURL(slug)
	if err != nil {
		if err == redis.Nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"error": "Url not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func InsertURLController(c *fiber.Ctx) error {
	body := &insertURLResponseBody{}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	
	errors := validateInsertURL(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"errors": errors,
		})
	}

	url, err := services.InsertURL(body.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
		
	}

	return c.Status(fiber.StatusOK).JSON(url)
}