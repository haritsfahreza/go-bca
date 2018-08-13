package example

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/haritsfahreza/go-bca"
	"github.com/haritsfahreza/go-bca/business"
)

func main() {
	client := business.NewClient(bca.Config{
		ClientID:     "",
		ClientSecret: "",
		APIKey:       "",
		APISecret:    "",
		CorporateID:  "BCAAPI2016", //Based on API document
		OriginHost:   "localhost",
	})

	ctx := context.Background()
	getBalanceInfo(ctx, client)
	getAccountStatement(ctx, client)
	fundTransfer(ctx, client)
}

func getBalanceInfo(ctx context.Context, client business.Client) {
	response, err := client.GetBalanceInfo(ctx, []string{"0201245680", "0063001004"})
	if err != nil {
		panic(err)
	}
	if len(*response.AccountDetailDataFailed) > 0 {
		for i, account := range *response.AccountDetailDataFailed {
			fmt.Printf("%d - Error: %s - %s", i, account.English, account.Indonesian)
		}
		return
	}
	for i, account := range *response.AccountDetailDataSuccess {
		jsonStr, _ := json.Marshal(account)
		fmt.Printf("%d - Account: %s", i, jsonStr)
	}
}

func getAccountStatement(ctx context.Context, client business.Client) {
	startDate, err := time.Parse("2006-01-02", "2016-08-29")
	if err != nil {
		panic(err)
	}

	endDate, err := time.Parse("2006-01-02", "2016-09-01")
	if err != nil {
		panic(err)
	}

	response, err := client.GetAccountStatement(ctx, "0201245680", startDate, endDate)
	if err != nil {
		panic(err)
	}

	jsonStr, _ := json.Marshal(response)
	fmt.Printf("Statement: %s", jsonStr)
}

func fundTransfer(ctx context.Context, client business.Client) {
	response, err := client.FundTransfer(ctx, bca.FundTransferRequest{
		CorporateID:         "BCAAPI2016",
		SourceAccountNumber: "0201245680",
		TransactionID:       "00000001",
		TransactionDate:     "2016-01-30",
		ReferenceID:         "12345/PO/2016",
		CurrencyCode:        "IDR",
		Amount:              float64(100000.00),
		BeneficiaryAccountNumber: "0201245681",
		Remark1:                  "Transfer Test",
		Remark2:                  "Online Transfer",
	})

	if err != nil {
		panic(err)
	}

	jsonStr, _ := json.Marshal(response)
	fmt.Printf("Statement: %s", jsonStr)
}
