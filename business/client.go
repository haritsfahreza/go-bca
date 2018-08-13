package business

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	bca "github.com/haritsfahreza/go-bca"
)

//Client is used to invoke BCA Business Banking API
type Client struct {
	Client      bca.ClientImplementation
	CorporateID string
	AccessToken string
}

//NewClient is used to initialize new business.Client
func NewClient(config bca.Config) Client {
	return Client{
		CorporateID: config.CorporateID,
		Client:      bca.NewClient(config),
	}
}

//GetBalanceInfo is used to get account balance information using account number(s)
func (c Client) GetBalanceInfo(ctx context.Context, accountNumbers []string) (bca.BalanceInfoResponse, error) {
	var response bca.BalanceInfoResponse
	path := fmt.Sprintf("/banking/v3/corporates/%s/accounts/%s", c.CorporateID, strings.Join(accountNumbers, ","))
	if err := c.Client.Call("GET", path, c.AccessToken, bytes.NewBufferString(""), response); err != nil {
		return response, err
	}
	return response, nil
}

//GetAccountStatement is used to get account statement information using account number
func (c Client) GetAccountStatement(ctx context.Context, accountNumber string, startDate, endDate time.Time) (bca.AccountStatementResponse, error) {
	var response bca.AccountStatementResponse
	path := fmt.Sprintf("/banking/v3/corporates/%s/accounts/%s/statements", c.CorporateID, accountNumber)

	v := url.Values{}
	v.Add("StartDate", startDate.Format("2006-01-02"))
	v.Add("EndDate", endDate.Format("2006-01-02"))
	path += "?" + v.Encode()

	if err := c.Client.Call("GET", path, c.AccessToken, bytes.NewBufferString(""), response); err != nil {
		return response, err
	}
	return response, nil
}

//FundTransfer is used to send fund transfer request
func (c Client) FundTransfer(ctx context.Context, request bca.FundTransferRequest) (bca.FundTransferResponse, error) {
	var response bca.FundTransferResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/banking/corporates/transfers"
	if err := c.Client.Call("POST", path, c.AccessToken, bytes.NewBuffer(jsonReq), response); err != nil {
		return response, err
	}
	return response, nil
}
