package main

import (
	"app/user/controller"
	"app/user/infrastructure/config"
	"app/user/infrastructure/repository"
	"app/user/usecase"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Hoge(c *gin.Context) string {
	return "hoge"
}
func main() {
	dbConn, err := sql.Open(config.SqlDriver, config.DatabaseUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		userController.Index(c)
	})
	router.Run()
}
