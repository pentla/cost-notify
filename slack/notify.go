package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pentla/cost-notify/entity"
	"golang.org/x/xerrors"
)

type SlackText struct {
	Text string `json:"text"`
	// Username   string `json:"username"`
	// Icon_emoji string `json:"icon_emoji"`
	// Icon_url   string `json:"icon_url"`
	// Channel    string `json:"channel"`
}

func PostBudget(webhookURL string, budget *entity.DailyBudget) error {
	payload := SlackText{
		Text: fmt.Sprintf("GCPの予算の%d %%に達しました。現在の課金額: %.2f, 予算: %.2f", budget.AlertThresholdExceeded, budget.CostAmount, budget.BudgetAmount),
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return xerrors.Errorf("marshal error: %v", err)
	}
	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		return xerrors.Errorf("Request error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return xerrors.Errorf("Response statuscode is not 200: %s %v", resp.Status, err)
		}
		return xerrors.Errorf("Response statuscode is not 200: %s %s", resp.Status, string(bodyBytes))
	}
	return nil
}
