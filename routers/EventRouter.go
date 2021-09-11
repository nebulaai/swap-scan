package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "swap-scan/common"
	"swap-scan/common/constants"
	"swap-scan/logs"
	"swap-scan/models"
)

func EventLogManager(router *gin.RouterGroup) {
	router.GET("/logs/:contractAddress/:cid", GetEventLogData)
}

func GetEventLogData(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	cid := c.Param("cid")
	eventListMap := make(map[string]interface{})
	goerliEventList, err := models.FindEventGoerlis(&models.EventGoerli{ContractAddress: contractAddress, PayloadCid: cid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(constants.GET_EVENT_FROM_DB_ERROR_CODE, constants.GET_EVENT_FROM_DB_ERROR_MSG))
		return
	}
	eventListMap[constants.NETWORK_TYPE_GOERLI] = goerliEventList
	polygonEventList, err := models.FindEventPolygons(&models.EventPolygon{ContractAddress: contractAddress, PayloadCid: cid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(constants.GET_EVENT_FROM_DB_ERROR_CODE, constants.GET_EVENT_FROM_DB_ERROR_MSG))
		return
	}
	eventListMap[constants.NETWORK_TYPE_POLYGON] = polygonEventList
	c.JSON(http.StatusOK, common.CreateSuccessResponse(eventListMap))
}
