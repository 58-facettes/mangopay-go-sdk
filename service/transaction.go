package service

import (
	"encoding/json"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

func parseTransaction(data []byte) (*model.Transaction, error) {
	var res model.Transaction
	return &res, json.Unmarshal(data, &res)
}

func parseTransactionCollection(data []byte) ([]model.Transaction, error) {
	var res []model.Transaction
	return res, json.Unmarshal(data, &res)
}
