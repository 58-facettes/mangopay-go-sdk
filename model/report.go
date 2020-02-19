package model

import (
	"github.com/58-facettes/mangopay-go-sdk/model/currency"
)

// Report gives the possibility to download huge lists of transactions or wallets to CSV format
// for accounting or analysis purposes. This can be done either from the API or Dashboard.
type Report struct {
	// The date when the report was executed
	ReportDate int64 `json:"ReportDate"`
	// The URL to download the report
	DownloadURL string `json:"DownloadURL"`
	// A URL that we will ping when the report is ready to download (works in a similar way to the hooks)
	CallbackURL string `json:"CallbackURL"`
	// The format of the report download
	DownloadFormat DownloadFormat `json:"DownloadFormat"`
	// The type of report
	ReportType ReportType `json:"ReportType"`
	// The column to sort against and direction
	Sort string `json:"Sort"`
	// Whether the report should be limited to the first 10 lines (and therefore quicker to execute)
	Preview bool `json:"Preview"`
	// An object of various filters for the report
	Filters ReportFilter `json:"Filters"`
	// A []string of columns/infos to show in the report
	// - you can choose from:
	// 		Alias, AuthorId, BankAccountId, BankWireRef, CardId, CardType, Country, CreationDate,
	// 		CreationDate:ISO,CreditedFundsAmount, CreditedFundsCurrency, CreditedUserId, CreditedWalletId,
	// 		Culture, DebitedFundsAmount, DebitedFundsCurrency, DebitedWalletId, DeclaredDebitedFundsAmount,
	// 		DeclaredDebitedFundsCurrency, DeclaredFeesAmount, DeclaredFeesCurrency, ExecutionDate,
	// 		ExecutionDate:ISO ExecutionType, ExpirationDate,ExpirationDate:ISO, FeesAmount, FeesCurrency,
	// 		Id, Nature, PaymentType, PreauthorizationId, ResultCode, ResultMessage, Status, Tag, Type,
	// 		WireReference
	Columns []string `json:"Columns"`
	// The result code
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode
	ResultMessage string `json:"ResultMessage"`
}

// DownloadFormat is the file format for the report download
type DownloadFormat string

const (
	DownloadFormatCSV DownloadFormat = "CSV"
)

// ReportType is the type of report
type ReportType string

const (
	ReportTransaction ReportType = "TRANSACTIONS"
)

type ReportFilter struct {
	// To return only resources that have CreationDate BEFORE this date
	BeforeDate int64 `json:"BeforeDate,omitempty"`
	// To return only resources that have CreationDate AFTER this date
	AfterDate int64 `json:"AfterDate,omitempty"`
	// The type of the transaction
	Type []string `json:"Type,omitempty"`
	// The status of the transaction
	Status []string `json:"Status,omitempty"`
	// The nature of the transaction
	Nature []string `json:"Nature,omitempty"`
	// The minimum amount of DebitedFunds
	MinDebitedFundsAmount int `json:"MinDebitedFundsAmount,omitempty"`
	// The currency for the minimum amount of DebitedFunds
	MinDebitedFundsCurrency currency.ISO3 `json:"MinDebitedFundsCurrency,omitempty"`
	// The maximum amount of DebitedFunds
	MaxDebitedFundsAmount int `json:"MaxDebitedFundsAmount,omitempty"`
	// The currency for the maximum amount of DebitedFunds
	MaxDebitedFundsCurrency currency.ISO3 `json:"MaxDebitedFundsCurrency,omitempty"`
	// The minimum amount of Fees
	MinFeesAmount int `json:"MinFeesAmount,omitempty"`
	// The currency for the minimum amount of Fees
	MinFeesCurrency currency.ISO3 `json:"MinFeesCurrency,omitempty"`
	// The maximum amount of Fees
	MaxFeesAmount int `json:"MaxFeesAmount,omitempty"`
	// The currency for the maximum amount of Fees
	MaxFeesCurrency currency.ISO3 `json:"MaxFeesCurrency,omitempty"`
	// A user's ID
	AuthorID string `json:"AuthorId,omitempty"`
	// The ID of a wallet
	WalletID string `json:"WalletId,omitempty"`
}

type ReportTransactionCreate struct {
	// A URL that we will ping when the report is ready to download (works in a similar way to the hooks)
	CallbackURL string `json:"CallbackURL,omitempty"`
	// The format of the report download
	DownloadFormat DownloadFormat `json:"DownloadFormat,omitempty"`
	// The column to sort against and direction
	Sort string `json:"Sort,omitempty"`
	// Whether the report should be limited to the first 10 lines (and therefore quicker to execute)
	Preview bool `json:"Preview,omitempty"`
	// An object of various filters for the report
	Filters ReportFilter `json:"Filters,omitempty"`
	// A list of columns/infos to show in the report - you can choose from:
	//		Alias, AuthorId, BankAccountId, BankWireRef, CardId, CardType, Country, CreationDate,CreationDate:ISO,
	//		DebitedFundsAmount, CreditedFundsAmount, CreditedFundsCurrency, CreditedUserId, CreditedWalletId, Culture,
	// 		DebitedFundsCurrency, DebitedWalletId, DeclaredDebitedFundsAmount, DeclaredDebitedFundsCurrency,
	// 		DeclaredFeesAmount, DeclaredFeesCurrency, ExecutionDate,ExecutionDate:ISO ExecutionType,
	// 		ExpirationDate,ExpirationDate:ISO, FeesAmount, FeesCurrency, Id, Nature, PaymentType, PreauthorizationId,
	// 		ResultCode, ResultMessage, Status, Tag, Type, WireReference
	Columns []string `json:"Columns,omitempty"`
}

type ReportWalletCreate struct {
	// A URL that we will ping when the report is ready to download (works in a similar way to the hooks)
	CallbackURL string `json:"CallbackURL,omitempty"`
	// The format of the report download
	DownloadFormat DownloadFormat `json:"DownloadFormat,omitempty"`
	// The column to sort against and direction
	Sort string `json:"Sort,omitempty"`
	// Whether the report should be limited to the first 10 lines (and therefore quicker to execute)
	Preview bool `json:"Preview,omitempty"`
	// An object of various filters for the report
	Filters ReportFilter `json:"Filters,omitempty"`
	// A list of columns/infos to show in the report - you can choose from: Id, Tag, CreationDate,
	// Owners, Description, BalanceAmount, BalanceCurrency, Currency, FundsType
	Columns []string `json:"Columns,omitempty"`
}
