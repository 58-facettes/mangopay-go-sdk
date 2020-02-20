package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Reports is responsible of all Reporting services.
type Reports struct{}

// TransactionCreate is creating a report for a Transaction.
func (Reports) TransactionCreate(payload *model.ReportTransactionCreate) (res *model.Report, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "reports/transactions/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// WalletCreate is creating a report for a Wallet.
func (Reports) WalletCreate(payload *model.ReportWalletCreate) (res *model.Report, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "reports/wallets/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retreve a Report from it's ID.
func (Reports) View(reportID string) (res *model.Report, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "reports/"+reportID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// List retreve all reports.
func (Reports) List(query *model.Query) (res []model.Report, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("reports/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
