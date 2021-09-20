package ethclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"swap-scan/config"
	"swap-scan/logs"
	"time"
)

type ConnSetup struct {
	ConnWeb *ethclient.Client
}

//setup polygon client connection
var WebConn = new(ConnSetup)

func ClientInit() {

	for {
		rpcUrl := config.GetConfig().NbaiMainnetNode.RpcUrl
		client, err := ethclient.Dial(rpcUrl)
		if err != nil {
			logs.GetLogger().Error("Try to reconnect block chain node" + rpcUrl + " ...")
			time.Sleep(time.Second * 5)
		} else {
			WebConn.ConnWeb = client
			break
		}
	}
}
