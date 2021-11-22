package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/pentla/cost-notify/cloudfunction"
)

func main() {
	ctx := context.Background()
	var m pubsub.Message
	m.Data = []byte(`{
		"budgetDisplayName": "name-of-budget",
		"alertThresholdExceeded": 1.0,
		"costAmount": 100.01,
		"costIntervalStart": "2019-01-01T00:00:00Z",
		"budgetAmount": 100.00,
		"budgetAmountType": "SPECIFIED_AMOUNT",
		"currencyCode": "USD"
	}`)
	err := cloudfunction.CostNotify(ctx, m)
	if err != nil {
		fmt.Println(err)
	}
}
