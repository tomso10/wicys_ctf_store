package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/config"
)

func logout(ctx *gin.Context) {
	ctx.SetCookie("auth", "", -1, "/", config.Hostname, false, false)
	ctx.Redirect(http.StatusFound, "/")
}
