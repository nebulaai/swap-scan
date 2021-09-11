package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"strings"
	"time"
)

type Configuration struct {
	Port            string          `toml:"port"`
	Dev             bool            `toml:"dev"`
	Database        database        `toml:"data_base"`
	NbaiMainnetNode NbaiMainnetNode `toml:"nbai_mainnet_node"`
	BscMainnetNode  BscMainnetNode  `toml:"bsc_mainnet_node"`
	ScheduleRule    ScheduleRule    `toml:"schedule_rule"`
}

type database struct {
	DbHost       string `toml:"db_host"`
	DbPort       string `toml:"db_port"`
	DbSchemaName string `toml:"db_schema_name"`
	DbUsername   string `toml:"db_username"`
	DbPwd        string `toml:"db_pwd"`
	DbArgs       string `toml:"db_args"`
}

type NbaiMainnetNode struct {
	RpcUrl                    string        `toml:"rpc_url"`
	PaymentContractAddress    string        `toml:"payment_contract_address"`
	ContractFunctionSignature string        `toml:"contract_function_signature"`
	ScanStep                  int64         `toml:"scan_step"`
	StartFromBlockNo          int64         `toml:"start_from_blockNo"`
	CycleTimeInterval         time.Duration `toml:"cycle_time_interval"`
}

type BscMainnetNode struct {
	RpcUrl                          string `toml:"rpc_url"`
	BscAdminWallet                  string `toml:"bsc_admin_wallet"`
	ChildChainManageContractAddress string `toml:"child_chain_manage_contract_address"`
	GasLimit                        uint64 `toml:"gas_limit"`
	ChainID                         int64  `toml:"chain_ID"`
}

type ScheduleRule struct {
	Nbai2BscMappingRedoRule string `toml:"nbai2bsc_mapping_redoRule"`
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

		{"data_base", "db_host"},
		{"data_base", "db_port"},
		{"data_base", "db_schema_name"},
		{"data_base", "db_username"},
		{"data_base", "db_pwd"},

		{"nbai_mainnet_node", "rpc_url"},
		{"nbai_mainnet_node", "payment_contract_address"},
		{"nbai_mainnet_node", "contract_function_signature"},
		{"nbai_mainnet_node", "scan_step"},
		{"nbai_mainnet_node", "cycle_time_interval"},

		{"bsc_mainnet_node", "rpc_url"},
		{"bsc_mainnet_node", "child_chain_manage_contract_address"},
		{"bsc_mainnet_node", "gas_limit"},
		{"bsc_mainnet_node", "chain_ID"},

		{"schedule_rule", "nbai2bsc_mapping_redoRule"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
