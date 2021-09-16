package models

import (
	"swap-scan/common/constants"
	"swap-scan/database"
)

type SwapCoinTransaction struct {
	ID          int64  `json:"id"`
	TxHashFrom  string `json:"tx_hash_from"`
	TxHashTo    string `json:"tx_hash_to"`
	Status      string `json:"status"`
	BlockNo     uint64 `json:"block_no"` //in nbai
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	GasFeeUsed  string `json:"gas_fee_used"`
	Quantity    string `json:"quantity"`
	BlockNoBsc  uint64 `json:"block_no_bsc"`
	FromNetwork int64  `json:"from_network"`
	ToNetwork   int64  `json:"to_network"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"create_at"`
	RedoTimes   int8   `json:"redo_times"`
}

// FindChildChainTransaction (&ChildChainTransaction{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindChildChainTransaction(whereCondition interface{}, orderCondition, limit, offset string) ([]*SwapCoinTransaction, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*SwapCoinTransaction
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
