package bca

import (
	"time"
)

//AccountBalance represents account balance information
type AccountBalance struct {
	AccountNumber    string
	Currency         string  `json:",omitempty"`
	Balance          float64 `json:",string"`
	AvailableBalance float64 `json:",string"`
	FloatAmount      float64 `json:",string"`
	HoldAmount       float64 `json:",string"`
	Plafon           float64 `json:",string"`
	Indonesian       string  `json:",omitempty"`
	English          string  `json:",omitempty"`
}

//BalanceInfoResponse represents account balance information response message
type BalanceInfoResponse struct {
	Error
	AccountDetailDataSuccess []AccountBalance `json:",omitempty"`
	AccountDetailDataFailed  []AccountBalance `json:",omitempty"`
}

//AccountStatement represents account statement information
type AccountStatement struct {
	TransactionDate   string
	BranchCode        string
	TransactionType   string
	TransactionAmount float64 `json:",string"`
	TransactionName   string
	Trailer           string
}

//AccountStatementResponse represents account statement response message
type AccountStatementResponse struct {
	Error
	StartDate    time.Time `json:",string"`
	EndDate      time.Time `json:",string"`
	Currency     string
	StartBalance float64 `json:",string"`
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
	Amount                   float64 `json:",string"`
	BeneficiaryAccountNumber string
	Remark1                  string
	Remark2                  string
}

//FundTransferResponse represents fund transfer response message
type FundTransferResponse struct {
	Error
	TransactionID   string
	TransactionDate string
	ReferenceID     string
	Status          string
}
