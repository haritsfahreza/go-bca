package bca

import (
	"time"
)

//AccountBalance represents account balance information
type AccountBalance struct {
	AccountNumber    string
	Currency         string  `json:",omitempty"`
	Balance          float64 `json:",omitempty"`
	AvailableBalance float64 `json:",omitempty"`
	FloatAmount      float64 `json:",omitempty"`
	HoldAmount       float64 `json:",omitempty"`
	Plafon           float64 `json:",omitempty"`
	Indonesian       string  `json:",omitempty"`
	English          string  `json:",omitempty"`
}

//BalanceInfoResponse represents account balance information response message
type BalanceInfoResponse struct {
	AccountDetailDataSuccess *[]AccountBalance `json:",omitempty"`
	AccountDetailDataFailed  *[]AccountBalance `json:",omitempty"`
}

//AccountStatement represents account statement information
type AccountStatement struct {
	TransactionDate   string
	BranchCode        string
	TransactionType   string
	TransactionAmount float64
	TransactionName   string
	Trailer           string
}

//AccountStatementResponse represents account statement response message
type AccountStatementResponse struct {
	StartDate    time.Time
	EndDate      time.Time
	Currency     string
	StartBalance float64
	Data         []AccountStatement
}

//FundTransferRequest represents fund transfer request message
type FundTransferRequest struct {
	CorporateID              string
	SourceAccountNumber      string
	TransactionID            string
	TransactionDate          string
	ReferenceID              string
	CurrencyCode             string
	Amount                   float64
	BeneficiaryAccountNumber string
	Remark1                  string
	Remark2                  string
}

//FundTransferResponse represents fund transfer response message
type FundTransferResponse struct {
	TransactionID   string
	TransactionDate string
	ReferenceID     string
	Status          string
}
