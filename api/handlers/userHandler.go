package handlers

import (
	"net/http"

	"api/auth"
	"api/config"
	"api/dao"
	"api/r"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (u *UserHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user dao.User
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, r.Error[string]("invalid request"))
		}

		if user.Name != config.USER_NAME || user.Password != config.PASSWORD {
			ctx.JSON(http.StatusUnauthorized, r.Error[string]("invalid user"))
		}

		token, err := auth.GenToken(user.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, r.Error[string]("generate token failed"))
		}

		ctx.JSON(http.StatusOK, r.Ok(token))
	}
}
