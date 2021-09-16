package models

import (
	"errors"
	"swap-scan/common/constants"
	"swap-scan/database"
	"swap-scan/logs"
)

type NetworkInfo struct {
	ID          int64  `json:"id"`
	Uuid        string `json:"uuid"`
	NetworkName string `json:"network_name"`
	CoinName    string `json:"coin_name"`
}

// FindNetworkInfo (&NetworkInfo{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindNetworkInfo(whereCondition interface{}, orderCondition, limit, offset string) ([]*NetworkInfo, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*NetworkInfo
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

func GetNetworkInfoByUUID(uuid string) (*NetworkInfo, error) {
	networkList, err := FindNetworkInfo(&NetworkInfo{Uuid: uuid}, "id desc", "10", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if len(networkList) > 0 {
		return networkList[0], nil
	}
	err = errors.New("According to the uuid = " + uuid + " ,no data found")
	return nil, err
}
