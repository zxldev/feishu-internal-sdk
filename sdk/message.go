package sdk

import (
	"github.com/sirupsen/logrus"
	"github.com/zxldev/feishu-internal-sdk/core/consts"
	"github.com/zxldev/feishu-internal-sdk/core/model"
)

func (t Tenant) SendMessage(msg model.Msg) {
	r, err := t.Post(consts.ApiRobotSendMessage, msg)
	if err != nil {
		logrus.Info(err.Error())
	} else {
		logrus.Info(string(r))
	}
}
