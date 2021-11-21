package entity

import "strconv"

type DailyBudget struct {
	CostAmount             float64
	AlertThresholdExceeded int
	BudgetAmount           float64
}

/*
Cost Alertのpayload
{
    "budgetDisplayName": "name-of-budget",
    "alertThresholdExceeded": 1.0,
    "costAmount": 100.01,
    "costIntervalStart": "2019-01-01T00:00:00Z",
    "budgetAmount": 100.00,
    "budgetAmountType": "SPECIFIED_AMOUNT",
    "currencyCode": "USD"
}
*/

func ParseDailyBudget(data map[string]string) (*DailyBudget, error) {
	costAmount, err := strconv.ParseFloat(data["costAmount"], 64)
	if err != nil {
		return nil, err
	}
	alertThresholdExceeded, err := strconv.ParseFloat(data["alertThresholdExceeded"], 64)
	if err != nil {
		return nil, err
	}
	alertExceedPercent := int(alertThresholdExceeded * 10)
	budgetAmount, err := strconv.ParseFloat(data["budgetAmount"], 64)
	if err != nil {
		return nil, err
	}
	return &DailyBudget{
		CostAmount:             costAmount,
		AlertThresholdExceeded: alertExceedPercent,
		BudgetAmount:           budgetAmount,
	}, nil
}
