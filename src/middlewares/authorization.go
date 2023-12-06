package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/ai-interview/repository"
)

func CheckJwtToken(c *fiber.Ctx) error {
	// リクエストヘッダーからAuthorizationの取得
	cookie := c.Cookies("IMP")
	user := repository.User{}
	loginUser, err := user.FindByCookie(cookie)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
	}

	if len(loginUser) == 0 {
		return c.Status(fiber.StatusForbidden).SendString("Session Tokenが不正です")
	}

	// LocalsメソッドでMiddlewareで取得した値をコンテキストに設定できる
	c.Locals("myUserId", loginUser[0].ID)

	// LocalsメソッドでMiddlewareで取得した値をコンテキストから取得する
	// token := c.Locals("bearer") // これで取得できる
	return c.Next()
}
