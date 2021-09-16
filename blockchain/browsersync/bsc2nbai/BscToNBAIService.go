package bsc2nbai

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

func swapNbaiFromBscToNbai(eventLog types.Log, txHashInNbai string, blockNo uint64, childChainTractionID int64) error {
	pk := os.Getenv(constants.PRIVATE_KEY_NAME_FOR_NBAI_ADMIN_WALLET)
	fromAddress := common.HexToAddress(config.GetConfig().BscAdminWallet)
	client := nbaiclient.WebConn.ConnWeb
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
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
	privateKey, _ := crypto.HexToECDSA(pk)
	callOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetConfig().NbaiMainnetNode.ChainID))

	//callOpts := new(bind.TransactOpts)
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = config.GetConfig().NbaiToBsc.GasLimit
	callOpts.Context = context.Background()

	/*childManagerAddress := common.HexToAddress(config.GetConfig().BscToNbai.BscSwapToNbaiContractAddress)
	childInstance, _ := goBind.NewChildChainManagerContract(childManagerAddress, client)*/

	rootManagerAddress := common.HexToAddress(config.GetConfig().BscToNbai.BscSwapToNbaiContractAddress)
	rootInstance, _ := goBind.NewRootChainManagerContract(rootManagerAddress, client)

	childChainTX := new(models.ChildChainTransaction)
	childChainTX.Status = constants.TRANSACTION_STATUS_FAIL

	var buff bytes.Buffer
	err = eventLog.EncodeRLP(&buff)
	if err != nil {
		logs.GetLogger().Error(err)
	}

	tx, err := rootInstance.Exit(callOpts, buff.Bytes())
	if err != nil {
		logs.GetLogger().Error(err)
		childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
		childChainTX.TxHashInBsc = ""
		childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
	} else {
		childChainTX.TxHashInBsc = tx.Hash().Hex()
		childChainTX.Status = constants.TRANSACTION_STATUS_SUCCESS
	}

	if err == nil {
		txMsg, err := tx.AsMessage(types.NewEIP155Signer(big.NewInt(config.GetConfig().NbaiMainnetNode.ChainID)), nil)
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
			}
			childChainTX.GasFeeUsed = strconv.FormatUint(txRecept.GasUsed, 10)
		}
	}
	quantity := new(big.Int)
	quantity.SetBytes(eventLog.Data)
	childChainTX.Quantity = quantity.String()

	if len(txRecept.Logs) > 0 {
		eventLogs := txRecept.Logs

		childChainTX.BlockNoBsc = eventLogs[0].BlockNumber
	}

	if childChainTractionID > 0 {
		childChainTX.ID = childChainTractionID
	}

	childChainTX.TxHashInNbai = txHashInNbai
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

func getContractBalance(client *ethclient.Client, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		logs.GetLogger().Error(err)
	}
	fmt.Println(weiToEther(balance))
	return balance, err
}

func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func EncodeToBytes(p interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return buf.Bytes(), nil
}
