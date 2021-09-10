package models

import (
	"swap-scan/common/constants"
	"swap-scan/database"
)

type ChildChainTransaction struct {
	ID           int64  `json:"id"`
	TxHashInBsc  string `json:"tx_hash_in_bsc"`
	TxHashInNbai string `json:"tx_hash_in_nbai""`
	Status       string `json:"status"`
	BlockNo      uint64 `json:"block_no"` //in nbai
	FromAddress  string `json:"from_address"`
	ToAddress    string `json:"to_address"`
	GasFeeUsed   string `json:"gas_fee_used"`
	Quantity     string `json:"quantity"`
	BlockNoBsc   uint64 `json:"block_no_bsc"`
	CreateAt     string `json:"create_at"`
	UpdateAt     string `json:"create_at"`
	RedoTimes    int8   `json:"redo_times"`
}

// FindChildChainTransaction (&ChildChainTransaction{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindChildChainTransaction(whereCondition interface{}, orderCondition, limit, offset string) ([]*ChildChainTransaction, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*ChildChainTransaction
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
