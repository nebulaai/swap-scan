package nbai2bsc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"strconv"
	"strings"
	"swap-scan/blockchain/initclient/bscclient"
	"swap-scan/common/constants"
	"swap-scan/common/utils"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	"swap-scan/models"
	"swap-scan/on-chain/goBind"
)

func ChangeNbaiToBnb(data []byte, txHashInNbai string, blockNo uint64, childChainTractionID int64) error {
	pk := os.Getenv(constants.PRIVATE_KEY_NAME_FOR_BSC_ADMIN_WALLET)
	fromAddress := common.HexToAddress(config.GetConfig().BscAdminWallet)
	client := bscclient.WebConn.ConnWeb
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
	callOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetConfig().BscMainnetNode.ChainID))

	//callOpts := new(bind.TransactOpts)
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = config.GetConfig().NbaiToBsc.GasLimit
	callOpts.Context = context.Background()

	childManagerAddress := common.HexToAddress(config.GetConfig().NbaiToBsc.NbaiSwapTobscContractAddress)
	childInstance, _ := goBind.NewChildChainManagerContract(childManagerAddress, client)

	childChainTX := new(models.SwapCoinTransaction)
	childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
	tx, err := childInstance.OnStateReceive(callOpts, big.NewInt(int64(1)), data)
	if err != nil {
		logs.GetLogger().Error(err)
		childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
		childChainTX.TxHashTo = ""
		childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
	} else {
		childChainTX.TxHashTo = tx.Hash().Hex()
		childChainTX.Status = constants.TRANSACTION_STATUS_SUCCESS
	}
	if err == nil {
		txMsg, err := tx.AsMessage(types.NewEIP155Signer(big.NewInt(config.GetConfig().BscMainnetNode.ChainID)), nil)
		if err != nil {
			logs.GetLogger().Error(err)
		}
		if err == nil {
			childChainTX.ToAddress = tx.To().Hex()
			childChainTX.FromAddress = txMsg.From().Hex()
		}
		txRecept, err := utils.CheckTx(bscclient.WebConn.ConnWeb, tx)
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
		if len(txRecept.Logs) > 0 {
			eventLogs := txRecept.Logs
			/*userWallet := hex.EncodeToString(eventLogs[0].Topics[1].Bytes())
			childChainTX.UserNbaiWallet = userWallet*/
			quantity := new(big.Int)
			quantity.SetBytes(eventLogs[0].Data)
			childChainTX.Quantity = quantity.String()
			childChainTX.BlockNoBsc = eventLogs[0].BlockNumber
		}
	}

	if childChainTractionID > 0 {
		childChainTX.ID = childChainTractionID
	}

	fromNetwork, _ := models.GetNetworkInfoByUUID(constants.NETWORK_INFO_UUID_FOR_NBAI)
	childChainTX.FromNetwork = fromNetwork.ID
	toNetwork, _ := models.GetNetworkInfoByUUID(constants.NETWORK_INFO_UUID_FOR_BSC)
	childChainTX.ToNetwork = toNetwork.ID
	childChainTX.TxHashFrom = txHashInNbai
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
