package router

import (
	"os"
	"rest-api-udemy/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	t := e.Group("/tasks")
	// NOTE: Useを使うことでエンドポイントにミドルウェアを追加できる
	t.Use(echojwt.WithConfig(echojwt.Config{
		// NOTE: トークンの署名キーを生成したときのキーを指定
		SigningKey:  []byte(os.Getenv("SECRET")),
		// NOTE: トークンの格納場所を指定(今回はcookieに対してtokenという名前で格納したのでそれを指定)
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	return e
}
