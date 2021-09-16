package nbai2bsc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"strings"
	"swap-scan/blockchain/initclient/nbaiclient"
	"swap-scan/common/utils"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	"swap-scan/models"
	"swap-scan/on-chain/goBind"
	"time"
)

/**
 * created on 08/20/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in swap-scan/LICENSE
 */

// EventLogSave Find the event that executed the contract and save to db
func ScanNbaiEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("nbai blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.StateSenderABI)
	paymentAbiString, err := abi.JSON(strings.NewReader(string(goBind.StateSenderABI)))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(config.GetConfig().NbaiToBsc.NbaiToBscEventTransferContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := config.GetConfig().NbaiToBsc.NbaiToBscEventContractEventFunctionSignature

	//test block no. is : 5297224
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	var logsInChain []types.Log
	var flag bool = true
	for flag {
		logsInChain, err = nbaiclient.WebConn.ConnWeb.FilterLogs(context.Background(), query)
		if err != nil {
			logs.GetLogger().Error(err)
			time.Sleep(5 * time.Second)
			continue
		}
		if err == nil {
			flag = false
		}
	}

	for _, vLog := range logsInChain {
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			eventList, err := models.FindEventNbai(&models.EventNbai{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				receiveMap := map[string]interface{}{}
				err = paymentAbiString.UnpackIntoMap(receiveMap, "StateSynced", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				var event = new(models.EventNbai)
				event.BlockNo = vLog.BlockNumber
				event.TxHash = vLog.TxHash.Hex()
				event.ContractName = "SwanPayment"
				event.ContractAddress = contractAddress.String()
				event.BytesData = receiveMap["data"].([]byte)
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				tx, _, err := nbaiclient.WebConn.ConnWeb.TransactionByHash(context.Background(), vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.ToAddress = tx.To().Hex()
					txMsg, err := tx.AsMessage(types.NewEIP155Signer(big.NewInt(999)), nil)
					if err != nil {
						logs.GetLogger().Error(err)
					} else {
						event.FromAddress = txMsg.From().Hex()
					}
				}

				txInLog, _, _ := nbaiclient.WebConn.ConnWeb.TransactionByHash(context.Background(), vLog.TxHash)
				if txInLog != nil {
					quantity := new(big.Int)
					quantity.SetBytes(txInLog.Value().Bytes())
					event.Quantity = quantity.String()
				}

				err = database.SaveOne(event)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				logs.GetLogger().Info("*************************nbai to bsc swaping start************************** ")
				err = ChangeNbaiToBnb(receiveMap["data"].([]byte), vLog.TxHash.Hex(), vLog.BlockNumber, 0)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				logs.GetLogger().Info("*************************nbai to bsc swaping end************************** ")
			}
		}
	}
	return nil
}
