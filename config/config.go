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
	BscAdminWallet  string          `toml:"bsc_admin_wallet"`
	NbaiAdminWallet string          `toml:"nbai_admin_wallet"`
	Database        database        `toml:"data_base"`
	NbaiMainnetNode NbaiMainnetNode `toml:"nbai_mainnet_node"`
	BscMainnetNode  BscMainnetNode  `toml:"bsc_mainnet_node"`
	EthMainnetNode  EthMainnetNode  `toml:"eth_mainnet_node"`
	ScheduleRule    ScheduleRule    `toml:"schedule_rule"`
	NbaiToBsc       NbaiToBsc       `toml:"nbai_to_bsc"`
	BscToNbai       BscToNbai       `toml:"bsc_to_nbai"`
	EthToBsc        EthToBsc        `toml:"eth_to_bsc"`
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
	RpcUrl  string `toml:"rpc_url"`
	ChainID int64  `toml:"chain_ID"`
}

type BscMainnetNode struct {
	RpcUrl  string `toml:"rpc_url"`
	ChainID int64  `toml:"chain_ID"`
}

type EthMainnetNode struct {
	RpcUrl  string `toml:"rpc_url"`
	ChainID int64  `toml:"chain_ID"`
}

type NbaiToBsc struct {
	NbaiToBscEventContractAddress                string        `toml:"nbai_to_bsc_event_contract_address"` // for scan event
	NbaiToBscEventContractEventFunctionSignature string        `toml:"nbai_to_bsc_event_contract_event_function_signature"`
	ScanStep                                     int64         `toml:"scan_step"`
	StartFromBlockNo                             int64         `toml:"start_from_blockNo"`
	CycleTimeInterval                            time.Duration `toml:"cycle_time_interval"`
	NbaiSwapTobscContractAddress                 string        `toml:"nbai_swap_to_bsc_contract_address"`
	GasLimit                                     uint64        `toml:"gas_limit"`
}

type BscToNbai struct {
	StartFromBlockNo                             int64         `toml:"start_from_blockNo"`
	CycleTimeInterval                            time.Duration `toml:"cycle_time_interval"`
	ScanStep                                     int64         `toml:"scan_step"`
	BscToNbaiEventContractAddress                string        `toml:"bsc_to_nbai_event_contract_address"` // for scan event
	BscToNbaiEventContractEventFunctionSignature string        `toml:"bsc_to_nbai_event_contract_event_function_signature"`
	BscSwapToNbaiContractAddress                 string        `toml:"bsc_swap_to_nbai_contract_address"`
}

type EthToBsc struct {
	StartFromBlockNo                            int64         `toml:"start_from_blockNo"`
	CycleTimeInterval                           time.Duration `toml:"cycle_time_interval"`
	ScanStep                                    int64         `toml:"scan_step"`
	EthToBscEventContractAddress                string        `toml:"eth_to_bsc_event_contract_address"` // for scan event
	EthToBscEventContractEventFunctionSignature string        `toml:"eth_to_bsc_event_contract_event_function_signature"`
	EthSwapToBscContractAddress                 string        `toml:"eth_swap_to_bsc_contract_address"`
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
		{"nbai_mainnet_node", "chain_ID"},

		{"bsc_mainnet_node", "rpc_url"},
		{"bsc_mainnet_node", "chain_ID"},

		{"nbai_to_bsc", "nbai_to_bsc_event_contract_address"},
		{"nbai_to_bsc", "nbai_to_bsc_event_contract_event_function_signature"},
		{"nbai_to_bsc", "scan_step"},

		{"bsc_to_nbai", "bsc_to_nbai_event_contract_address"},
		{"bsc_to_nbai", "bsc_to_nbai_event_contract_event_function_signature"},
		{"bsc_to_nbai", "start_from_blockNo"},
		{"bsc_to_nbai", "cycle_time_interval"},
		{"bsc_to_nbai", "scan_step"},

		{"schedule_rule", "nbai2bsc_mapping_redoRule"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
