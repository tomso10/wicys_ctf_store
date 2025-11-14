package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "pong"})
}
