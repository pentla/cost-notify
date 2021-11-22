package entity

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

type CostMessage struct {
	BudgetDisplayName      string  `json:"budgetDisplayName"`
	AlertThresholdExceeded float64 `json:"alertThresholdExceeded"`
	CostAmount             float64 `json:"costAmount"`
	CostIntervalStart      string  `json:"costIntervalStart"`
	BudgetAmount           float64 `json:"budgetAmount"`
	BudgetAmountType       string  `json:"budgetAmountType"`
	CurrencyCode           string  `json:"currencyCode"`
}

func ParseDailyBudget(data CostMessage) (*DailyBudget, error) {
	alertExceedPercent := int(data.AlertThresholdExceeded * 10)
	return &DailyBudget{
		CostAmount:             data.CostAmount,
		AlertThresholdExceeded: alertExceedPercent,
		BudgetAmount:           data.BudgetAmount,
	}, nil
}
