package cloudfunction

import (
	"context"
	"encoding/json"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/pentla/cost-notify/entity"
	"github.com/pentla/cost-notify/slack"
	"golang.org/x/xerrors"
)

type SlackPayload struct{}

var webhook = os.Getenv("SLACK_WEBHOOK")

func CostNotify(ctx context.Context, m pubsub.Message) error {
	var data entity.CostMessage
	err := json.Unmarshal(m.Data, &data)
	if err != nil {
		return xerrors.Errorf("json marshal error: %v", err)
	}
	budget, err := entity.ParseDailyBudget(data)
	if err != nil {
		return xerrors.Errorf("Failed to parse message: %v", err)
	}
	err = slack.PostBudget(webhook, budget)
	if err != nil {
		return xerrors.Errorf("Failed to post slack: %v", err)
	}
	return nil
}
