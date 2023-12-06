package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview/repository"
)

func InterviewSaveHandler(c *fiber.Ctx) error {
	var interview repository.Interview
	if err := c.BodyParser(&interview); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// 保存
	if err := interview.Save(); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(interview)
}

func GetInterviewHandler(c *fiber.Ctx) error {
	var interview repository.Interview
	records, err := interview.FindById(1)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(records)
}
