package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/bot"
	_ "gitlab.ritsec.cloud/competitions/ists-2023/store/config"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/requests"
)

var (
	Router *gin.Engine
)

func init() {
	// gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	Router.SetTrustedProxies(nil)
	Router.LoadHTMLGlob("templates/*.tmpl")
	Router.Static("/assets", "./assets")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nShutting Down....")
		cleanup()
		os.Exit(1)
	}()
}

func main() {
	requests.LoadRequests(Router)
	bot.Init()

	Router.Run(":80")
}

func cleanup() {
	data.Close()
	bot.Close()
}
