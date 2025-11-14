package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/user"
)

func index(ctx *gin.Context) {
	_user, err := auth.Parse(ctx)
	if err != nil && err != auth.ErrUnauthorized {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	} else if err == auth.ErrUnauthorized {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
		})
		return
	}

	if itemName, ok := ctx.GetQuery("itemName"); ok {
		items, _ := data.Item.GetAll()
		if _user.Permissions != user.PermissionsBlack {
			ctx.HTML(http.StatusOK, "index.tmpl",
				gin.H{
					"title":  "MILITARY PORTAL",
					"items":  items,
					"user":   _user,
					"navbar": true,
					"info":   fmt.Sprintf("Successfully purchased %s", itemName),
				},
			)
		} else {
			ctx.HTML(http.StatusOK, "index.tmpl",
				gin.H{
					"title":       "MILITARY PORTAL",
					"items":       items,
					"user":        _user,
					"navbar":      true,
					"info":        fmt.Sprintf("Successfully purchased %s", itemName),
					"navbar_link": "/admin",
				},
			)
		}
		return
	}

	if _token, ok := ctx.GetQuery("token"); ok {
		items, _ := data.Item.GetAll()
		if _user.Permissions != user.PermissionsBlack {
			ctx.HTML(http.StatusOK, "index.tmpl",
				gin.H{
					"title":  "MILITARY PORTAL",
					"items":  items,
					"user":   _user,
					"navbar": true,
					"info":   fmt.Sprintf("Successfully redeemed %s", _token),
				},
			)
		} else {
			ctx.HTML(http.StatusOK, "index.tmpl",
				gin.H{
					"title":       "MILITARY PORTAL",
					"items":       items,
					"user":        _user,
					"navbar":      true,
					"info":        fmt.Sprintf("Successfully redeemed %s", _token),
					"navbar_link": "/admin",
				},
			)
		}
		return
	}

	if _user.Permissions != user.PermissionsBlack {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   _user,
				"navbar": true,
			},
		)
		return
	}

	items, _ := data.Item.GetAll()

	ctx.HTML(http.StatusOK, "index.tmpl",
		gin.H{
			"title":       "MILITARY PORTAL",
			"items":       items,
			"user":        _user,
			"navbar":      true,
			"navbar_link": "/admin",
		},
	)
}
