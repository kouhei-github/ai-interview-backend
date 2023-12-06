package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kouhei-github/ai-interview/route"
)

func main() {
	app := fiber.New()

	// ルーターの設定
	router := &route.Router{FiberApp: app}

	// CORS (Cross Origin Resource Sharing)の設定
	// アクセスを許可するドメイン等を設定します
	router.FiberApp.Use(cors.New(cors.Config{AllowHeaders: "Origin, Content-Type, Accept"}))

	route.LoadRouter(router)

	// Webサーバー起動時のエラーハンドリング => localhostの時コメントイン必要
	if err := router.FiberApp.Listen(":80"); err != nil {
		panic(err)
	}
}
