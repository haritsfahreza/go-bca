# 🏦 BCA (Bank Central Asia) API's Go Library

[![Library Status](https://img.shields.io/badge/status-unofficial-yellow.svg)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/haritsfahreza/go-bca)](https://goreportcard.com/report/github.com/haritsfahreza/go-bca)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)
[![Build Status](https://travis-ci.org/haritsfahreza/go-bca.svg?branch=master)](https://travis-ci.org/haritsfahreza/go-bca)

Go(lang) library to speed up your BCA (Bank Central Asia) API integration process. See this [official documentation of BCA API](https://developer.bca.co.id/documentation/)

## Usage
```
import (
	"context"

	"github.com/haritsfahreza/go-bca"
	"github.com/haritsfahreza/go-bca/business"
)

func main() {
    cfg := bca.Config{
		URL:          "https://sandbox.bca.co.id",
		ClientID:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		ClientSecret: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		APIKey:       "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		APISecret:    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		CorporateID:  "BCAAPI2016", //Based on API document
		OriginHost:   "localhost",
	}
	businessClient := business.NewClient(cfg)
	authClient := auth.NewClient(cfg)

	ctx := context.Background()
	authToken, err := authClient.GetToken(ctx)
	if err != nil {
		panic(err)
	}

    businessClient.AccessToken = authToken.AccessToken

    ctx := context.Background()
    response, err := client.GetBalanceInfo(ctx, []string{"0201245680", "0063001004"})
}
```

## Example

We have attached usage examples in this repository in folder `example`.
Please proceed there for more detail on how to run the example.

## License

See [LICENSE](LICENSE).