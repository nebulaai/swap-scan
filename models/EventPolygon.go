package models

import (
	"swap-scan/common/constants"
	"swap-scan/database"
)

type EventPolygon struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
	IndexedData     string `json:"indexed_data"`
	ContractName    string `json:"contract_name"`
	ContractAddress string `json:"contract_address"`
	LockedFee       string `json:"locked_fee"`
	Deadline        string `json:"deadline"`
	PayloadCid      string `json:"payload_cid"`
	BlockNo         uint64 `json:"block_no"`
	MinerAddress    string `json:"miner_address"`
}

func (self *EventPolygon) FindOneEventPolygon(condition interface{}) (*EventPolygon, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventPolygons(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventPolygon, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventPolygon
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
