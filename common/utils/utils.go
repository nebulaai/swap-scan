package utils

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"io/ioutil"
	"math/big"
	"os"
	"reflect"
	"swap-scan/logs"
	"time"
)

// GetEpochInMillis get current timestamp
func GetEpochInMillis() (millis int64) {
	nanos := time.Now().UnixNano()
	millis = nanos / 1000000
	return
}

func ReadContractAbiJsonFile(aptpath string) (string, error) {
	jsonFile, err := os.Open(aptpath)

	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	return string(byteValue), nil
}

func GetRewardPerBlock() *big.Int {
	rewardBig, _ := new(big.Int).SetString("35000000000000000000", 10) // the unit is wei
	return rewardBig
}

func CheckTx(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
retry:
	/*checkTImes--
	if checkTImes <= 0 {
		err := errors.New("check tx status time out! txhash=" + tx.Hash().Hex())
		return nil, err
	}*/
	rp, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		if err == ethereum.NotFound {
			logs.GetLogger().Error("tx %v not found, check it later ", tx.Hash().String())
			time.Sleep(1 * time.Second)
			goto retry
		} else {
			logs.GetLogger().Error("TransactionReceipt fail: %s", err)
			return nil, err
		}
	}
	return rp, nil
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}
	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
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

func AddZeroToTheFrontOfTheArray(arr []byte, zeroNo int) []byte {
	var zeroArr []byte
	for i := 0; i < zeroNo; i++ {
		zeroArr = append(zeroArr, 0)
	}
	arr = append(zeroArr, arr...)
	return arr
}
