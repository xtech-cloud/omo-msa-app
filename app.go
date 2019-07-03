package main

import (
	"os"

	"github.com/xtech-cloud/omo-msa-app/httpAPI"
	"github.com/xtech-cloud/omo-msa-app/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.AutoMigrateDatabase()

	httpAddrArg := os.Getenv("APP_HTTP_ADDR")
	if "" == httpAddrArg {
		httpAddrArg = ":80"
	}

	// --------------------
	// Http API
	// --------------------
	router := gin.Default()
	router.POST("/app/create", httpAPI.HandleCreateApp)
	router.POST("/app/query", httpAPI.HandleQueryApp)
	router.POST("/app/list", httpAPI.HandleListApp)
	router.POST("/app/profile/modify", httpAPI.HandleModifyAppProfile)
	router.POST("/app/secret/reset", httpAPI.HandleResetSecret)
	router.POST("/app/key/reset", httpAPI.HandleResetKey)

	router.Run(httpAddrArg)
}
