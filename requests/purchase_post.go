package requests

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/discord"
)

var (
	ErrItemNotFound  error = errors.New("item not found")
	ErrInvalidItemID error = errors.New("invalid item ID")
)

type purchaseS struct {
	Instructions string
	Name         string
}

func purchase_post(ctx *gin.Context) {
	_user, err := auth.Parse(ctx)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	_data := purchaseS{
		Instructions: ctx.PostForm("instructions"),
	}

	_data.Name = ctx.PostForm("item")
	if _data.Name == "" {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   _user,
				"navbar": true,
				"error":  ErrInvalidItemID.Error(),
			},
		)
		return
	}

	item, err := data.Item.Get(_data.Name)

	if err != nil {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   _user,
				"navbar": true,
				"error":  ErrItemNotFound.Error(),
			},
		)
		return
	}

	_, err = data.Transaction.Create(_user.Username, item.Name, _data.Instructions)
	if err != nil {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   _user,
				"navbar": true,
				"error":  err.Error(),
			},
		)
		return
	}

	discord.Send(_user.Username, item.Name, _data.Instructions)

	q := url.Values{}
	q.Set("itemName", item.Name)
	location := url.URL{Path: "/", RawQuery: q.Encode()}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
