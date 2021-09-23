package eth2bsc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"strings"
	"swap-scan/blockchain/initclient/ethclient"
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
func ScanEthEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	stateSendABI, err := abi.JSON(strings.NewReader(string(goBind.StateSenderABI)))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	//read contract api json file
	logs.GetLogger().Println("eth blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	contractAddress := common.HexToAddress(config.GetConfig().NbaiOnEthToBsc.NbaiOnEthToBscEventContractAddress)
	contractFunctionSignature := config.GetConfig().NbaiOnEthToBsc.NbaiOnEthToBscEventContractEventFunctionSignature

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	var logsInChain []types.Log
	var flag bool = true
	for flag {
		logsInChain, err = ethclient.WebConn.ConnWeb.FilterLogs(context.Background(), query)
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
			eventList, err := models.FindEventNbaiOnEth(&models.EventNbaiOnEth{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				receiveMap := map[string]interface{}{}
				err = stateSendABI.UnpackIntoMap(receiveMap, "StateSynced", vLog.Data)

				userWalletAddress := common.BytesToAddress(vLog.Data[32*5 : 32*6])
				nbaiERC20Address := common.BytesToAddress(vLog.Data[32*6 : 32*7])
				//depositeType := vLog.Data[32*2 : 32*3]

				swapNbaiERC20ValueInBytes := vLog.Data[32*9 : 32*10]
				swapNbaiERC20Value := new(big.Int)
				swapNbaiERC20Value.SetBytes(swapNbaiERC20ValueInBytes)

				var event = new(models.EventNbaiOnEth)
				event.FromAddress = userWalletAddress.Hex()
				event.BlockNo = vLog.BlockNumber
				event.TxHash = vLog.TxHash.Hex()
				event.ContractName = "ChildChainManagerProxy"
				event.ContractAddress = contractAddress.String()
				event.BytesData = vLog.Data
				event.NBAIERC20AddressOnEth = nbaiERC20Address.Hex()
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				tx, _, err := ethclient.WebConn.ConnWeb.TransactionByHash(context.Background(), vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.ToAddress = tx.To().Hex()
				}
				event.Quantity = swapNbaiERC20Value.String()
				err = database.SaveOne(event)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}

				haveSwapped, err := models.CheckHaveSwapSuccess(vLog.TxHash.Hex())
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				} else {
					if !haveSwapped {
						logs.GetLogger().Info("*************************nbaierc20 on eth to bsc swaping start************************** ")
						err = ChangNBAIERC20OnEthToBnb(receiveMap["data"].([]byte), vLog.TxHash.Hex(), vLog.BlockNumber, 0, swapNbaiERC20Value)
						if err != nil {
							logs.GetLogger().Error(err)
							logs.GetLogger().Info("*************************nbaierc20 on  eth to bsc swaping end************************** ")
							continue
						}
						logs.GetLogger().Info("*************************nbaierc20 on  eth to bsc swaping end************************** ")
					}
				}
			}
		}
	}
	return nil
}
