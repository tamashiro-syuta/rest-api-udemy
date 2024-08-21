package main

import (
	"rest-api-udemy/controller"
	"rest-api-udemy/db"
	"rest-api-udemy/repository"
	"rest-api-udemy/router"
	"rest-api-udemy/usecase"
)

func main() {
	// NOTE: DBのインスタンスを作成
	db := db.NewDB()
	// NOTE: DBのインスタンスを引数にしてリポジトリのコンストラクターを呼び出す
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	// NOTE: エラーが発生した場合はechoのロガーにエラーを出力して終了させる
	e.Logger.Fatal(e.Start(":8080"))
}
