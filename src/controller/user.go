package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview/repository"
	"github.com/kouhei-github/ai-interview/utils/password"
)

func SignUpHandler(c *fiber.Ctx) error {
	// リクエストボディの受け取り
	var user *repository.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// saltを生成
	user.Salt = password.RandomString(15)
	// パスワードのハッシュ化
	hashed, err := password.HashPassword(user.Password + user.Salt)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user.Password = hashed

	// 保存
	if err := user.Save(); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UserAllHandler(c *fiber.Ctx) error {
	localValue := c.Locals("myUserId")
	user := repository.User{}
	if err := user.FindById(localValue.(uint)); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
