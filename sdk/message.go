package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zxldev/feishu-internal-sdk/core/consts"
	"github.com/zxldev/feishu-internal-sdk/core/model"
	"net/http"
)

func (t Tenant) SendMessage(msg model.Msg) {
	r, err := t.Post(consts.ApiRobotSendMessage, msg)
	if err != nil {
		logrus.Info(err.Error())
	} else {
		logrus.Info(string(r))
	}
}

func (t Tenant) RobotMessageCallback(w http.ResponseWriter, r *http.Request, eventCallback func(event model.Event), buttonCallBack func(button model.ButtonCallback)) {
	rCopy := *r
	event := model.Event{}
	err := json.NewDecoder(rCopy.Body).Decode(&event)
	if event.Token != "" {
		w.WriteHeader(404)
		return
	}
	if err != nil && event.Type != "" {
		if event.Type == "url_verification" {
			r.Write(bytes.NewBufferString(`{"challenge":"` + event.Challenge + `"}`))
			return
		} else if event.Type == "event_callback" {
			eventCallback(event)
			return
		}
	}
	//TODO v2 的消息 https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM

	//开始按照button 回调处理

	buttonData := model.ButtonCallback{}
	json.NewDecoder(rCopy.Body).Decode(&buttonData)
	buttonCallBack(buttonData)

}
