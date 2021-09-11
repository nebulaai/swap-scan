package schedule

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"strings"
	"swap-scan/blockchain/browsersync/nbai"
	"swap-scan/blockchain/initclient/nbaiclient"
	"swap-scan/common/constants"
	"swap-scan/config"
	"swap-scan/logs"
	"swap-scan/models"
	"swap-scan/on-chain/goBind"
	"time"
)

func RedoMapping() error {
	txList, err := models.FindChildChainTransaction(&models.ChildChainTransaction{Status: constants.TRANSACTION_STATUS_FAIL, TxHashInBsc: ""}, "create_at desc", "-1", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, v := range txList {

		paymentAbiString, err := abi.JSON(strings.NewReader(string(goBind.StateSenderABI)))
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		//SwanPayment contract address
		contractAddress := common.HexToAddress(config.GetConfig().NbaiMainnetNode.PaymentContractAddress)

		//test block no. is : 5297224
		blockNoInt, err := strconv.Atoi(strconv.FormatUint(v.BlockNo, 10))
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(blockNoInt)),
			ToBlock:   big.NewInt(int64(blockNoInt)),
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
			if vLog.TxHash.Hex() == strings.TrimSpace(v.TxHashInNbai) {
				fmt.Println(vLog.BlockHash.Hex())
				fmt.Println(vLog.BlockNumber)
				fmt.Println(vLog.TxHash.Hex())
				receiveMap := map[string]interface{}{}
				err = paymentAbiString.UnpackIntoMap(receiveMap, "StateSynced", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}

				err = nbai.ChangeNbaiToBnb(receiveMap["data"].([]byte), v.TxHashInNbai, v.BlockNo, v.ID)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
			}
		}
	}

	return nil
}
