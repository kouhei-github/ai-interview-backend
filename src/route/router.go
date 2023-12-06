package route

import (
	"github.com/kouhei-github/ai-interview/controller"
	"github.com/kouhei-github/ai-interview/middlewares"
)

func (router *Router) GetRouter() {
	// 練習
	router.FiberApp.Get("/", controller.HelloHandler)
	// パスパラメータ取得
	router.FiberApp.Get("/path/:id", controller.PathParamTestHandler)
	// パスパラメータ取得
	router.FiberApp.Get("/query", controller.QueryParamTestHandler)

	// ヘルスチェック
	router.FiberApp.Post("/test", controller.HealthCheckHandler)

	router.FiberApp.Get("/realface/:id", controller.PathParamTestHandler)

	router.FiberApp.Post("/signup", controller.SignUpHandler)
	router.FiberApp.Post("/login", controller.LoginHandler)
	router.FiberApp.Get("/user", middlewares.CheckJwtToken, controller.UserAllHandler)

	// 応募者の追加
	router.FiberApp.Post("/applicant", controller.ApplicantSaveHandler)

	// 面接の追加
	router.FiberApp.Post("/interview", controller.InterviewSaveHandler)

	// 面接情報の追加
	router.FiberApp.Get("/interview", controller.GetInterviewHandler)

	// 応募者の取得
	router.FiberApp.Get("/applicant", controller.FindApplicantHandler)
}
