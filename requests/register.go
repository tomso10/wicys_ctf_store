package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
)

type registerS struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Team     string `json:"team"`
}

func register(ctx *gin.Context) {
	jsonBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	var _data registerS

	err = json.Unmarshal(jsonBody, &_data)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	_, err = data.User.Create(_data.Username, _data.Password, _data.Team)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "success"})

}
