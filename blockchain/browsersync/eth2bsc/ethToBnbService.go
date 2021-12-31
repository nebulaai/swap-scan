package eth2bsc

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

func ChangNBAIERC20OnEthToBnb(data []byte, txHashInNbai string, blockNo uint64, childChainTractionID int64, swapValue *big.Int) error {
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

	childManagerAddress := common.HexToAddress(config.GetConfig().NbaiOnEthToBsc.EthSwapToBscContractAddress)
	childInstance, _ := goBind.NewChildChainManagerContract(childManagerAddress, client)

	childChainTX := new(models.SwapCoinTransaction)
	childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
	if swapValue.Sign() > 0 {
		tx, err := childInstance.OnStateReceive(callOpts, big.NewInt(int64(1)), data)
		if err != nil {
			logs.GetLogger().Error(err)
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
			txReicept, err := utils.CheckTx(bscclient.WebConn.ConnWeb, tx)
			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				if txReicept.Status == uint64(1) {
					if childChainTX.FromAddress != "" {
						childChainTX.Status = constants.HTTP_STATUS_SUCCESS
						childChainTX.BlockNoBsc = txReicept.BlockNumber.Uint64()
						logs.GetLogger().Println("swap success! txHash=" + tx.Hash().Hex())
					} else {
						logs.GetLogger().Println("swap failed! txHash=" + tx.Hash().Hex())
					}
					childChainTX.GasFeeUsed = strconv.FormatUint(txReicept.GasUsed, 10)
				}
			}
		}
		childChainTX.Quantity = swapValue.String()
	} else {
		logs.GetLogger().Println("The balance is less than 100NBAI, no transfer operation has occurred")
		childChainTX.Status = constants.TRANSACTION_STATUS_SUCCESS
		childChainTX.TxHashTo = constants.TX_HASH_NIL_VALUE_FOR_CHARGED
		childChainTX.ToAddress = constants.WALLET_NIL_VALUE
		childChainTX.GasFeeUsed = "0"
		feeValue, _ := new(big.Int).SetString(config.GetConfig().NbaiOnEthToBsc.HandlingFee+constants.ZERO_18, 10)
		childChainTX.Quantity = swapValue.Add(swapValue, feeValue).String()
	}

	if childChainTractionID > 0 {
		childChainTX.ID = childChainTractionID
	}

	fromNetwork, _ := models.GetNetworkInfoByUUID(constants.NETWORK_INFO_UUID_FOR_ETH)
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
