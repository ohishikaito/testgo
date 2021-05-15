package controller

import (
	"app/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsercase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) *UserController {
	return &UserController{
		userUsercase: uu,
	}
}

func (u *UserController) Index(c *gin.Context) {
	// ここでcontext作るのだるいから、省略できんかなあ。毎回書かないと
	ctx := c.Request.Context()
	users, err := u.userUsercase.GetUsers(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": string(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
