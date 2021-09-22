package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	"swap-scan/blockchain/browsersync/eth2bsc"
	"swap-scan/blockchain/initclient/bscclient"
	"swap-scan/blockchain/initclient/ethclient"
	"swap-scan/blockchain/initclient/nbaiclient"
	"swap-scan/common/constants"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	"swap-scan/routers"
	"time"
)

func main() {
	LoadEnv()

	// init database
	db := database.Init()

	initMethod()

	eth2bsc.EthBlockBrowserSyncAndEventLogsSync()
	//go schedule.RedoMappingSchedule()

	//go bsc2nbai.BscBlockBrowserSyncAndEventLogsSync()

	//go nbai2bsc.NbaiBlockBrowserSyncAndEventLogsSync()

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := r.Group("/api/v1")
	routers.EventLogManager(v1.Group(constants.URL_EVENT_PREFIX))

	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}

func initMethod() string {
	config.InitConfig("")
	nbaiclient.ClientInit()
	bscclient.ClientInit()
	ethclient.ClientInit()
	return ""
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logs.GetLogger().Error(err)
	}
}
