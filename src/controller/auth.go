package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview/repository"
	"github.com/kouhei-github/ai-interview/utils/password"
	"os"
	"strconv"
)

func LoginHandler(c *fiber.Ctx) error {
	// リクエストボディの受け取り
	var user *repository.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// ログインユーザー取得
	loginUser, err := user.FindByEmail(user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// 存在するか
	exist := password.VerifyPassword(user.Password+loginUser[0].Salt, loginUser[0].Password)
	if !exist {
		fmt.Println("ユーザーが存在しません")
		return c.Status(fiber.StatusBadRequest).SendString("ユーザーが存在しません")
	}
	salt := password.RandomString(15)
	userId := strconv.FormatUint(uint64(loginUser[0].ID), 10)
	token, err := password.HashPassword(userId + salt)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	loginUser[0].SessionToken = token
	// 保存
	if err := loginUser[0].Update(); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:   "IMP",
		Value:  token,
		Path:   "/",
		Domain: os.Getenv("Domain"),
	})
	return c.Status(fiber.StatusOK).JSON(map[string]string{"token": token})
}
