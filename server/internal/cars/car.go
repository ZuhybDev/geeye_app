package cars

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
)

type NewCarParam struct {
	Name        string `json:"name"`
	Color       string `json:"color"`
	NumberPlate string `json:"number_plate"`
}

func (h *carHandler) CreateNewCar(c fiber.Ctx) error {

	var newCarInputs NewCarParam

	if err := c.Bind().Body(&newCarInputs); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if newCarInputs.Name == "" || newCarInputs.Color == "" || newCarInputs.NumberPlate == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "All feilds are required",
		})
	}

	params := db.NewCarParams{
		Name:        utils.ToPgTex(&newCarInputs.Name),
		Color:       utils.ToPgTex(&newCarInputs.Color),
		NumberPlate: utils.ToPgTex(&newCarInputs.NumberPlate),
	}

	car, err := h.app.Query.NewCar(c.Context(), params)

	if err != nil {
		fmt.Println("New Car: ", err)

		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to record new car try again",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Successfully recorded a new car",
		"car":     car,
	})

}

type UpdateCarInputs struct {
	Name        *string `json:"name"`
	Color       *string `json:"color"`
	NumberPlate *string `json:"number_plate"`
}

func (h *carHandler) UpdateCar(c fiber.Ctx) error {

	id := c.Params("id")

	parsedId, err := utils.ParsePGIDs(id)

	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse car id",
		})
	}

	var UpdateCarParam UpdateCarInputs

	if err := c.Bind().Body(&UpdateCarParam); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	params := db.UpdateCarParams{
		ID:          parsedId,
		Name:        utils.ToPgTex(UpdateCarParam.Name),
		Color:       utils.ToPgTex(UpdateCarParam.Color),
		NumberPlate: utils.ToPgTex(UpdateCarParam.NumberPlate),
	}

	car, err := h.app.Query.UpdateCar(c.Context(), params)

	if err != nil {
		fmt.Println("Update Car: ", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to update try again",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully updated",
		"car":     car,
	})

}

func (h *carHandler) DeleteCar(c fiber.Ctx) error {

	id := c.Params("id")

	parseId, err := utils.ParsePGIDs(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse car id",
		})
	}

	err = h.app.Query.DeleteCar(c.Context(), parseId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to delete please try again",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully deleted",
	})
}

func (h *carHandler) GetCarById(c fiber.Ctx) error {

	id := c.Params("id")

	parsedId, err := utils.ParsePGIDs(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse car id",
		})
	}

	car, err := h.app.Query.GetCarBYId(c.Context(), parsedId)

	return c.Status(200).JSON(fiber.Map{
		"car": car,
	})

}
