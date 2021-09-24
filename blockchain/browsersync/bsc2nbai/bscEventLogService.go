package bsc2nbai

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"swap-scan/blockchain/initclient/bscclient"
	"swap-scan/common/utils"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	"swap-scan/models"
	"time"
)

// EventLogSave Find the event that executed the contract and save to db
func ScanBSCEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("-----------------bsc blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.StateSenderABI)
	//paymentAbiString, err := abi.JSON(strings.NewReader(string(goBind.child)))
	/*childERC20AbiString, err := utils.ReadContractAbiJsonFile("on-chain/contracts/abi/ChildERC20.json")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	contractAbiInstans, err := abi.JSON(strings.NewReader(childERC20AbiString))*/

	//SwanPayment contract address
	contractAddress := common.HexToAddress(config.GetConfig().BscToNbai.BscToNbaiEventContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := config.GetConfig().BscToNbai.BscToNbaiEventContractEventFunctionSignature

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
	var err error
	for flag {
		logsInChain, err = bscclient.WebConn.ConnWeb.FilterLogs(context.Background(), query)
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
			if utils.IsZeroAddress(vLog.Topics[2].Hex()) {
				eventList, err := models.FindEventBsc(&models.EventBsc{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				if len(eventList) <= 0 {
					quantity := new(big.Int)
					quantity.SetBytes(vLog.Data)
					/*receiveMap := map[string]interface{}{}
					err = contractAbiInstans.UnpackIntoMap(receiveMap, "withdraw", vLog.Data)
					if err != nil {
						logs.GetLogger().Error(err)
						continue
					}*/
					var event = new(models.EventBsc)
					event.BlockNo = vLog.BlockNumber
					event.TxHash = vLog.TxHash.Hex()
					event.ContractName = "SwanPayment"
					event.ContractAddress = contractAddress.String()
					event.BytesData = vLog.Data
					event.Quantity = quantity.String()
					event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
					tx, _, err := bscclient.WebConn.ConnWeb.TransactionByHash(context.Background(), vLog.TxHash)
					if err != nil {
						logs.GetLogger().Error(err)
					} else {
						event.ToAddress = tx.To().Hex()
						txMsg, err := tx.AsMessage(types.NewEIP155Signer(big.NewInt(config.GetConfig().BscMainnetNode.ChainID)), nil)
						if err != nil {
							logs.GetLogger().Error(err)
						} else {
							event.FromAddress = txMsg.From().Hex()
						}
					}
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
							logs.GetLogger().Info("*************************bsc to nbai swaping start************************** ")
							err = swapNbaiFromBscToNbai(vLog, vLog.BlockNumber, 0)
							if err != nil {
								logs.GetLogger().Error(err)
								logs.GetLogger().Info("*************************bsc to nbai swaping end************************** ")
								continue
							}
							logs.GetLogger().Info("*************************bsc to nbai swaping end************************** ")
						}
					}
				}
			}

		}
	}
	return nil
}
