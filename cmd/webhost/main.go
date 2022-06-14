package main

import (
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/HalvaPovidlo/discordBotGo/cmd/config"
	"github.com/HalvaPovidlo/discordBotGo/pkg/zap"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	logger := zap.NewLogger(cfg.General.Debug)
	if !cfg.General.Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	serverIPHandler := func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Host.IP+":"+cfg.Host.Bot)
	}
	mockIPHandler := func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Host.IP+":"+cfg.Host.Mock)
	}

	router := gin.Default()
	router.Use(static.Serve("/web", static.LocalFile("./www/", true)))
	router.GET("/", func(c *gin.Context) {
		location := url.URL{Path: "/web"}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	})
	router.GET("/server", serverIPHandler)
	router.GET("/mock", mockIPHandler)
	go func() {
		err := router.Run(":" + cfg.Host.Web)
		if err != nil {
			logger.Error(err)
			return
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.Infow("Graceful shutdown")
	_ = logger.Sync()
}
