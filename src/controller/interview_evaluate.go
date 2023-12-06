package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview/repository"
	"strconv"
)

func SaveInterviewEvaluationHandler(c *fiber.Ctx) error {
	var interviewEvaluation repository.InterviewEvaluation
	if err := c.BodyParser(&interviewEvaluation); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := interviewEvaluation.Save(); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(interviewEvaluation)
}

func GetInterviewEvaluationHandler(c *fiber.Ctx) error {
	evaluationId := c.Params("id")
	evaluateId, err := strconv.Atoi(evaluationId)
	if err != nil {
		return err
	}

	var interviewEvaluation repository.InterviewEvaluation
	record, err := interviewEvaluation.FindById(uint(evaluateId))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(record)
}
