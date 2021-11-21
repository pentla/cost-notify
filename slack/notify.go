package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/pentla/cost-notify/entity"
	"golang.org/x/xerrors"
)

type SlackText struct {
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_emoji string `json:"icon_emoji"`
	Icon_url   string `json:"icon_url"`
	Channel    string `json:"channel"`
}

func PostBudget(webhookURL string, budget *entity.DailyBudget) error {
	text := SlackText{
		Text: fmt.Sprintf("GCPの予算の%d %%に達しました。現在の課金額: %f, 予算: %f", budget.AlertThresholdExceeded, budget.CostAmount, budget.BudgetAmount),
	}
	var values url.Values
	textBytes, _ := json.Marshal(text)
	err := json.Unmarshal(textBytes, &values)
	if err != nil {
		return xerrors.Errorf("Marshal error: %v", err)
	}
	resp, err := http.PostForm(webhookURL, values)
	if err != nil {
		return xerrors.Errorf("Request error: %v", err)
	}
	if resp.StatusCode != 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		return xerrors.Errorf("Response statuscode is not 200: %s %s", resp.Status, string(bodyBytes))
	}
	return nil
}
