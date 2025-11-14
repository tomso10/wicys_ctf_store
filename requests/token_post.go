package requests

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/auth"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
)

func token_post(ctx *gin.Context) {
	_user, err := auth.Parse(ctx)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
			"title":  "Sign In",
			"navbar": false,
			"error":  err.Error(),
		})
		return
	}

	token_form := ctx.PostForm("token")

	_token, err := data.Token.Redeem(_user.Username, token_form)
	if err != nil {
		items, _ := data.Item.GetAll()
		ctx.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title":  "MILITARY PORTAL",
				"items":  items,
				"user":   _user,
				"navbar": true,
				"error":  "Invalid Token",
			},
		)
		return
	}

	var _type string
	if len(_token.Edges.User) == 0 {
		_type = fmt.Sprintf("%v: %v", _token.Type, "unclaimed")
	} else {
		_type = fmt.Sprintf("%v: %v", _token.Type, "claimed before")
	}

	q := url.Values{}
	q.Set("token", fmt.Sprintf("%v (%v)", _token.Value, _type))
	location := url.URL{Path: "/", RawQuery: q.Encode()}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
