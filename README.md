# GCPCostNotify

GCPの予算アラートをSlackに送信するスクリプトです。

## 用意

1. GCPの予算アラートを作成します。
https://cloud.google.com/billing/docs/how-to/budgets?hl=ja

2. Cloud Pub/Subのトピックを作成し、予算とアラートから接続してください。

https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications?hl=ja

3. SlackのIncoming Webhookを作成します。

https://slack.com/intl/ja-jp/help/articles/115005265063-Slack-%E3%81%A7%E3%81%AE-Incoming-Webhook-%E3%81%AE%E5%88%A9%E7%94%A8

## 環境変数

- SLACK_WEBHOOK
    - slackの[incoming_webhook](https://slack.com/intl/ja-jp/help/articles/115005265063-Slack-%E3%81%A7%E3%81%AE-Incoming-Webhook-%E3%81%AE%E5%88%A9%E7%94%A8)で得られるURLです。

## デプロイ

```bash
gcloud functions deploy cost_notify --entry-point=CostNotify --runtime=go116 --triger-topic={Cloud Pub/Subのトピック名} --set-build-env-vars SLACK_WEBHOOK=xxx
```

詳しくは[CloudFunctionドキュメント](https://cloud.google.com/sdk/gcloud/reference/functions/deploy)より。

