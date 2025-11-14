package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
)

func purchase_get(ctx *gin.Context) {
	user, err := auth.Parse(ctx)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	itemName, ok := ctx.GetQuery("item")
	if !ok {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   user,
				"navbar": true,
			},
		)
		return
	}

	if !data.Item.Exists(itemName) {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   user,
				"navbar": true,
			},
		)
		return
	}

	item, _ := data.Item.Get(itemName)

	ctx.HTML(http.StatusOK, "purchase.tmpl", gin.H{
		"title":  "Purchase",
		"navbar": true,
		"user":   user,
		"item":   item,
	})
}
