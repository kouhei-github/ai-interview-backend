package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview/repository"
)

func ApplicantSaveHandler(c *fiber.Ctx) error {
	var applicant repository.Applicant
	if err := json.Unmarshal(c.Body(), &applicant); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// 保存
	if err := applicant.Save(); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(applicant)
}

func FindApplicantHandler(c *fiber.Ctx) error {
	var applicant repository.Applicant
	records, err := applicant.FindById(1)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(records)
}
