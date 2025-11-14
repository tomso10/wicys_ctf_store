package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
)

func login_get(ctx *gin.Context) {
	_, err := auth.Parse(ctx)
	if err != nil && err != auth.ErrUnauthorized {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title":  "Sign In",
		"navbar": false,
	})
}
