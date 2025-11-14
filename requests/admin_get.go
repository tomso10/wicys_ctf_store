package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/user"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/helpers"
)

func admin_get(ctx *gin.Context) {
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

	if _user.Permissions != user.PermissionsBlack  {
		ctx.HTML(http.StatusUnauthorized, "status.tmpl", gin.H{
			"navbar":             true,
			"user":               _user,
			"status_code":        "401",
			"status_description": "Unauthorized",
		})
	}

	ctx.HTML(http.StatusOK, "admin.tmpl",
		gin.H{
			"title":      "Admin Panel",
			"user":       _user,
			"navbar":     true,
			"admin_info": helpers.NewAdminInfo(),
		},
	)
}
