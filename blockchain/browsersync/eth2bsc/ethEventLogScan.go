package eth2bsc

import (
	"math/big"
	"strconv"
	"swap-scan/blockchain/initclient/ethclient"
	"swap-scan/blockchain/initclient/nbaiclient"
	"swap-scan/common/constants"
	"swap-scan/common/utils"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	models2 "swap-scan/models"
	"sync"
	"time"
)

func ethBlockBrowserSyncAndEventLogsSync() {
	startScanBlockNo := getStartBlockNo()

	for {
		var mutex sync.Mutex
		mutex.Lock()
		var blockNoCurrent *big.Int
		var err error
		var getBlockFlag bool = true
		for getBlockFlag {
			blockNoCurrent, err = nbaiclient.WebConn.GetBlockNumber()
			if err != nil {
				ethclient.ClientInit()
				logs.GetLogger().Error(err)
				time.Sleep(5 * time.Second)
				continue
			}
			if err == nil {
				getBlockFlag = false
			}
		}

		blockScanRecord := new(models2.BlockScanRecord)
		whereCondition := "network_type='" + constants.NETWORK_TYPE_ETH + "'"
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			startScanBlockNo = config.GetConfig().EthToBsc.StartFromBlockNo
		}
		if len(blockScanRecordList) > 0 {
			if blockScanRecordList[0].LastCurrentBlockNumber <= blockNoCurrent.Int64() {
				startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
			} else {
				startScanBlockNo = config.GetConfig().EthToBsc.StartFromBlockNo
			}
			blockScanRecord.ID = blockScanRecordList[0].ID
		}

		for {
			start := startScanBlockNo
			end := start + config.GetConfig().EthToBsc.ScanStep
			if startScanBlockNo > blockNoCurrent.Int64() {
				break
			}
			err = ScanEthEventFromChainAndSaveEventLogData(start, end)
			if err != nil {
				logs.GetLogger().Error(err)
				time.Sleep(time.Second * 1)
				continue
			}

			if end >= blockNoCurrent.Int64() {
				blockScanRecord.LastCurrentBlockNumber = blockNoCurrent.Int64()
			} else {
				blockScanRecord.LastCurrentBlockNumber = end
			}

			blockScanRecord.NetworkType = constants.NETWORK_TYPE_ETH
			blockScanRecord.UpdateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)

			err = database.SaveOne(blockScanRecord)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			start = end
			startScanBlockNo = end
			if end >= blockNoCurrent.Int64() {
				break
			}
		}

		getBlockFlag = true
		mutex.Unlock()

		time.Sleep(time.Second * config.GetConfig().EthToBsc.CycleTimeInterval)
		logs.GetLogger().Info("-------------------------eth----------------------------")
	}
}

func getStartBlockNo() int64 {
	var startScanBlockNo int64 = 1

	if config.GetConfig().EthToBsc.StartFromBlockNo > 0 {
		startScanBlockNo = config.GetConfig().EthToBsc.StartFromBlockNo
	}

	blockScanRecord := new(models2.BlockScanRecord)
	whereCondition := "network_type='" + constants.NETWORK_TYPE_ETH + "'"
	blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
	if err != nil {
		logs.GetLogger().Error(err)
		startScanBlockNo = config.GetConfig().EthToBsc.StartFromBlockNo
	}

	if len(blockScanRecordList) > 0 {
		if blockScanRecordList[0].LastCurrentBlockNumber > startScanBlockNo {
			startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
		}
	}
	return startScanBlockNo
}
