package models

import (
	"swap-scan/common/constants"
	"swap-scan/database"
)

type EventBsc struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
	IndexedData     string `json:"indexed_data"`
	ContractName    string `json:"contract_name"`
	ContractAddress string `json:"contract_address"`
	Quantity        string `json:"quantity"`
	BytesData       []byte `json:"bytes_data""`
	BlockNo         uint64 `json:"block_no"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	CreateAt        string `json:"create_at"`
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventBsc(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventBsc, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventBsc
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
