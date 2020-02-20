package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model/currency"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Transactions is responsible of all services for the Transaction.
type Transactions struct{}

// UserList is listing all transactions from a given user ID.
func (Transactions) UserList(userID string) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/transactions/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// WalletList is listing all transactions from a given Wallet ID.
func (Transactions) WalletList(walletID string) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "wallets/"+walletID+"/transactions/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// DisputesList is listing all transactions from a given Dispute ID.
func (Transactions) DisputesList(disputeID string) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "disputes/"+disputeID+"/transactions/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// ClientWalletList is listing all transactions from  client Wallets.
func (Transactions) ClientWalletList(fundType model.Funds, currency currency.ISO3) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/wallets/"+string(fundType)+"/"+string(currency)+"/transactions/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// PreAuthorizationList includes PreAuthorized PayIns related to the PreAuthorization.
// Indeed, if a PreAuthorized PayIn fails, you can re-use the same Preauthorization to create a new PayIn
// while the PreAuthorization has not expired. As soon as a PreAuthorized Payin has succeeded,
// you cannot use the PreAuthorization anymore - even if the amount was partial.
func (Transactions) PreAuthorizationList(preAuthorizationID string) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "preauthorizations/"+preAuthorizationID+"/transactions/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// BankAccountList is listing all transactions for a BankAccount.
// ?? in the documentation this is showed as a payload in body ??
// ?? I use it in the query params ??
// ?? end ??
//
// ResultCode string OPTIONAL
// The result code of the transaction (you can filter your transactions list by multiple ResultCode values, each one must be separated by a comma)
//
// Status TransactionStatus OPTIONAL
// The status of the transaction (you can filter your transactions list by multiple Status values, each one must be separated by a comma)
func (Transactions) BankAccountList(bankAccountID string, query *model.Query) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("bankaccounts/"+bankAccountID+"/transactions/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// CardList is listing all transactions for a given Card ID.
// ?? in the documentation this is showed as a payload in body ??
// ?? I use it in the query params ??
//
// ?? end ??
// Status TransactionStatus OPTIONAL
// The status of the transaction (you can filter your transactions list by multiple Status values, each one must be separated by a comma)
//
// ResultCode string OPTIONAL
// The result code of the transaction (you can filter your transactions list by multiple ResultCode values, each one must be separated by a comma)
func (Transactions) CardList(cardID string, query *model.Query) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("cards/"+cardID+"/transactions/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// MandateList is listing all transactions for a given Mandate ID.
// ?? in the documentation this is showed as a payload in body ??
// ?? I use it in the query params ??
// // ?? end ??
//
// Status TransactionStatus OPTIONAL
// The status of the transaction (you can filter your transactions list by multiple Status values, each one must be separated by a comma)
//
// ResultCode string OPTIONAL
// The result code of the transaction (you can filter your transactions list by multiple ResultCode values, each one must be separated by a comma)
func (Transactions) MandateList(mandateID string, query *model.Query) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("mandates/"+mandateID+"/transactions/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
