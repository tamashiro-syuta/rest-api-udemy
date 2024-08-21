package router

import (
	"net/http"
	"os"
	"rest-api-udemy/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	// NOTE: 全てのリクエストに対してCORSをチェックするミドルウェアを適用
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		// NOTE: cookieの送受信を許可
		AllowCredentials: true,
	}))
	// NOTE: CSRF対策のミドルウェアを適用
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, // NOTE: 自動でsecureがtrueになるので、デバッグ時はコメントアウトし、下のDefaultModeを有効にする
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

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
