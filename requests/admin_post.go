package requests

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/helpers"
)

type adminS struct {
	Username string
	Function string
	Amount   int
	UserID   int
}

var (
	ErrNotAValidInteger = errors.New("amount specified is not a valid integer")
)

func admin_post(ctx *gin.Context) {
	_user, err := auth.Parse(ctx)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "status.tmpl", gin.H{
			"navbar":             true,
			"user":               _user,
			"status_code":        "401",
			"status_description": "Unauthorized",
		})
		return
	}

	amount, err := strconv.Atoi(ctx.PostForm("amount"))
	if err != nil {
		ctx.HTML(http.StatusOK, "admin.tmpl",
			gin.H{
				"title":      "Admin Panel",
				"user":       _user,
				"navbar":     true,
				"admin_info": helpers.NewAdminInfo(),
				"error":      ErrNotAValidInteger,
			})
		return
	}

	_data := adminS{
		Username: ctx.PostForm("account"),
		Function: ctx.PostForm("function"),
		Amount:   amount,
	}

	data.User.Get(_data.Username)

	switch _data.Function {
	case "ADD":
		_, err = data.User.IncrementBalance(_data.Username, _data.Amount)
	case "SUB":
		_, err = data.User.DecrementBalance(_data.Username, _data.Amount)
	case "SET":
		_, err = data.User.SetBalance(_data.Username, _data.Amount)
	}

	if err != nil {
		ctx.HTML(http.StatusOK, "admin.tmpl",
			gin.H{
				"title":      "Admin Panel",
				"user":       _user,
				"navbar":     true,
				"admin_info": helpers.NewAdminInfo(),
				"error":      err,
			})
		return
	}

	ctx.HTML(http.StatusOK, "admin.tmpl",
		gin.H{
			"title":      "Admin Panel",
			"user":       _user,
			"navbar":     true,
			"admin_info": helpers.NewAdminInfo(),
		})
}
