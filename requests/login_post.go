package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/config"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
)

type loginS struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login_post(ctx *gin.Context) {
	_data := loginS{
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
	}

	_user, err := data.User.Authenticate(_data.Username, _data.Password)
	if err != nil {
		ctx.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	jwtString, expirationTime, err := auth.GenerateJWT(_user.Username)
	if err != nil {
		ctx.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title":  "Log In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.SetCookie("auth", jwtString, expirationTime, "/", config.Hostname, false, false)
	ctx.Redirect(http.StatusFound, "/")
}
