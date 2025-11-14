package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
)

func info(ctx *gin.Context) {
	_user, err := auth.Parse(ctx)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"username": _user.Username,
			"team":     _user.Permissions,
			"balance":  _user.Balance,
		},
	)
}
