package bsc2nbai

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"strconv"
	"strings"
	"swap-scan/blockchain/initclient/nbaiclient"
	"swap-scan/common/constants"
	"swap-scan/common/utils"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	"swap-scan/models"
	"swap-scan/on-chain/goBind"
)

func swapNbaiFromBscToNbai(eventLog types.Log, blockNo uint64, childChainTractionID int64) error {
	pk := os.Getenv(constants.PRIVATE_KEY_NAME_FOR_NBAI_ADMIN_WALLET)
	nbaiAdminWalletAddress := common.HexToAddress(config.GetConfig().NbaiAdminWallet)
	client := nbaiclient.WebConn.ConnWeb
	nonce, err := client.PendingNonceAt(context.Background(), nbaiAdminWalletAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	if strings.HasPrefix(strings.ToLower(pk), "0x") {
		pk = pk[2:]
	}
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	callOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = config.GetConfig().NbaiToBsc.GasLimit
	callOpts.Context = context.Background()

	rootManagerAddress := common.HexToAddress(config.GetConfig().BscToNbai.BscSwapToNbaiContractAddress)
	rootInstance, _ := goBind.NewRootChainManagerContract(rootManagerAddress, client)

	childChainTX := new(models.SwapCoinTransaction)
	childChainTX.Status = constants.TRANSACTION_STATUS_FAIL

	quantity := new(big.Int)
	quantity.SetBytes(eventLog.Data)
	childChainTX.Quantity = quantity.String()
	feeValue, _ := new(big.Int).SetString(config.GetConfig().BscToNbai.HandlingFee+constants.ZERO_18, 10)

	if quantity.Sub(quantity, feeValue).Sign() > 0 {

		fmt.Println(eventLog.Data)
		fmt.Println(len(eventLog.Data))
		zeroArray := utils.AddZeroToTheFrontOfTheArray(quantity.Bytes(), len(eventLog.Data)-len(quantity.Bytes()))
		eventLog.Data = zeroArray
		fmt.Println(eventLog.Data)
		fmt.Println(len(eventLog.Data))
		var buff bytes.Buffer
		err = eventLog.EncodeRLP(&buff)
		if err != nil {
			logs.GetLogger().Error(err)
		}
		fmt.Println(eventLog.Data)
		tx, err := rootInstance.Exit(callOpts, buff.Bytes())
		if err != nil {
			logs.GetLogger().Error(err)
			childChainTX.TxHashTo = ""
			childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
		} else {
			childChainTX.TxHashTo = tx.Hash().Hex()
			childChainTX.Status = constants.TRANSACTION_STATUS_SUCCESS
		}

		if err == nil {
			txMsg, err := tx.AsMessage(types.NewEIP155Signer(chainId), nil)
			if err != nil {
				logs.GetLogger().Error(err)
			}
			if err == nil {
				childChainTX.ToAddress = tx.To().Hex()
				childChainTX.FromAddress = txMsg.From().Hex()
			}
		}
		txRecept, err := utils.CheckTx(client, tx)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if txRecept.Status == uint64(1) {
				if childChainTX.FromAddress != "" {
					childChainTX.Status = constants.HTTP_STATUS_SUCCESS
					logs.GetLogger().Println("swap success! txHash=" + tx.Hash().Hex())
				} else {
					logs.GetLogger().Println("swap failed! txHash=" + tx.Hash().Hex())
				}
				childChainTX.GasFeeUsed = strconv.FormatUint(txRecept.GasUsed, 10)
			}
		}
		if len(txRecept.Logs) > 0 {
			eventLogs := txRecept.Logs
			childChainTX.BlockNo = eventLogs[0].BlockNumber
		}
	} else {
		logs.GetLogger().Println("The balance is less than 100NBAI, no transfer operation has occurred")
		childChainTX.Status = constants.HTTP_STATUS_SUCCESS
		childChainTX.ToAddress = constants.WALLET_NIL_VALUE
		childChainTX.TxHashTo = constants.TX_HASH_NIL_VALUE_FOR_NO_CHARGE
		childChainTX.FromAddress = constants.WALLET_NIL_VALUE
		childChainTX.GasFeeUsed = "0"
		childChainTX.Quantity = "0"
	}

	if childChainTractionID > 0 {
		childChainTX.ID = childChainTractionID
	}

	childChainTX.BlockNoBsc = eventLog.BlockNumber
	fromNetwork, _ := models.GetNetworkInfoByUUID(constants.NETWORK_INFO_UUID_FOR_BSC)
	childChainTX.FromNetwork = fromNetwork.ID
	toNetwork, _ := models.GetNetworkInfoByUUID(constants.NETWORK_INFO_UUID_FOR_NBAI)
	childChainTX.ToNetwork = toNetwork.ID

	childChainTX.TxHashFrom = eventLog.TxHash.Hex()
	childChainTX.BlockNo = blockNo
	currenTime := utils.GetEpochInMillis()
	childChainTX.CreateAt = strconv.FormatInt(currenTime, 10)
	childChainTX.UpdateAt = strconv.FormatInt(currenTime, 10)
	err = database.SaveOne(childChainTX)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
