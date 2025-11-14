package requests

import "github.com/gin-gonic/gin"

func LoadRequests(router *gin.Engine) {
	// router.GET("/ping", ping)
	// router.GET("/info", info)
	router.GET("/", index)
	router.GET("/index", index)
	router.GET("/login", login_get)
	router.GET("/logout", logout)
	router.GET("/admin", admin_get)
	router.GET("/purchase", purchase_get)

	// router.POST("/register", register)
	router.POST("/login", login_post)
	router.POST("/admin", admin_post)
	router.POST("/purchase", purchase_post)
	router.POST("/token", token_post)
}
