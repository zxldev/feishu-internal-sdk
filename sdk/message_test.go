package sdk

import (
	"github.com/zxldev/feishu-internal-sdk/core/model"
	"testing"
)

func TestSend(t *testing.T) {
	user := "zxldev@live.com"
	FeiShu.Init("", "")
	FeiShu.Tenant.SendMessage(model.Msg{
		Email:   user,
		MsgType: "interactive",
		Card: &model.Card{
			Header: &model.CardHeader{
				Title: &model.CardHeaderTitle{
					Tag:     "plain_text",
					Content: "xxxx标题",
				},
			},
			Elements: []interface{}{
				model.CardElementAction{
					Tag: "div",
					Text: &model.CardElementText{
						Tag:     "plain_text",
						Content: "申请人：" + user,
					},
				},
				model.CardElementAction{
					Tag: "div",
					Text: &model.CardElementText{
						Tag:     "plain_text",
						Content: "开始时间：立即生效",
					},
				},
				model.CardElementAction{
					Tag: "div",
					Text: &model.CardElementText{
						Tag:     "plain_text",
						Content: "结束时间：2小时后",
					},
				},
				model.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						model.ActionButton{
							Tag:  "button",
							Type: "default",
							Text: model.CardElementText{
								Tag:     "plain_text",
								Content: "通过",
							},
							Value: map[string]interface{}{
								"type":  "device_rw",
								"op":    "approval",
								"user":  user,
								"start": "now",
								"end":   "1h",
							},
						},
						model.ActionButton{
							Tag:  "button",
							Type: "default",
							Text: model.CardElementText{
								Tag:     "plain_text",
								Content: "不通过",
							},
							Value: map[string]interface{}{
								"type":  "device_rw",
								"op":    "cancel",
								"user":  user,
								"start": "now",
								"end":   "1h",
							},
						},
					},
				},
			},
		},
	})
}
