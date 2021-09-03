package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"strings"
	"time"
)

type Configuration struct {
	Port            string
	Database        database
	NbaiMainnetNode NbaiMainnetNode
	BscMainnetNode  BscMainnetNode
	ScheduleRule    ScheduleRule
	Dev             bool
}

type database struct {
	DbUsername   string
	DbPwd        string
	DbHost       string
	DbPort       string
	DbSchemaName string
	DbArgs       string
}

type NbaiMainnetNode struct {
	RpcUrl                    string
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
	CycleTimeInterval         time.Duration
}

type BscMainnetNode struct {
	RpcUrl                          string
	BscAdminWallet                  string
	ChildChainManageContractAddress string
	GasLimit                        uint64
	ChainID                         int64
	PaymentContractAddress          string
	ContractFunctionSignature       string
	ScanStep                        int64
	StartFromBlockNo                int64
	CycleTimeInterval               time.Duration
}

type ScheduleRule struct {
	Nbai2BscMappingRedoRule string
}

var config *Configuration

func InitConfig(configFile string) {
	if strings.Trim(configFile, " ") == "" {
		configFile = "./config/config.toml"
	}
	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func GetConfig() Configuration {
	if config == nil {
		InitConfig("")
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},

		{"DataBase", "dbHost"},
		{"DataBase", "dbPort"},
		{"DataBase", "dbSchemaName"},
		{"DataBase", "dbUsername"},
		{"DataBase", "dbPwd"},

		{"NbaiMainnetNode", "rpcUrl"},
		{"NbaiMainnetNode", "paymentContractAddress"},
		{"NbaiMainnetNode", "contractFunctionSignature"},
		{"NbaiMainnetNode", "scanStep"},
		{"NbaiMainnetNode", "cycleTimeInterval"},

		{"BscMainnetNode", "rpcUrl"},
		{"BscMainnetNode", "bscAdminWallet"},
		{"BscMainnetNode", "childChainManageContractAddress"},
		{"BscMainnetNode", "gasLimit"},
		{"BscMainnetNode", "chainID"},

		{"ScheduleRule", "nbai2BscMappingRedoRule"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
