package swaputil

import (
	"swap-scan/common/constants"
	"swap-scan/config"
	"swap-scan/logs"
	models2 "swap-scan/models"
)

func GetStartBlockNo(networkName string) int64 {
	var startScanBlockNo int64 = 1
	var whereCondition string

	switch networkName {
	case constants.NETWORK_TYPE_NBAI:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_NBAI + "'"
		if config.GetConfig().NbaiToBsc.StartFromBlockNo > 0 {
			startScanBlockNo = config.GetConfig().NbaiToBsc.StartFromBlockNo
		}
	case constants.NETWORK_TYPE_BSC:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_BSC + "'"
		if config.GetConfig().BscToNbai.StartFromBlockNo > 0 {
			startScanBlockNo = config.GetConfig().BscToNbai.StartFromBlockNo
		}
	case constants.NETWORK_TYPE_ETH:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_ETH + "'"
		if config.GetConfig().NbaiOnEthToBsc.StartFromBlockNo > 0 {
			startScanBlockNo = config.GetConfig().NbaiOnEthToBsc.StartFromBlockNo
		}
	}

	blockScanRecord := new(models2.BlockScanRecord)
	blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
	if err != nil {
		logs.GetLogger().Error(err)
		startScanBlockNo = config.GetConfig().NbaiToBsc.StartFromBlockNo
	}

	if len(blockScanRecordList) > 0 {
		if blockScanRecordList[0].LastCurrentBlockNumber > startScanBlockNo {
			startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
		}
	}
	return startScanBlockNo
}
