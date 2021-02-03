package sdk

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zxldev/feishu-internal-sdk/core/consts"
	"github.com/zxldev/feishu-internal-sdk/core/model"
	"io/ioutil"
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

func (t Tenant) RobotMessageCallback(w http.ResponseWriter, r *http.Request, eventCallback func(event model.Event), buttonCallBack func(button model.ButtonCallback)) (err error) {
	event := model.Event{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	if FeiShu.EncryptKey != "" { //设置了加密方式
		encodeMessage := model.EncodeMessage{}
		err = json.Unmarshal(body, &encodeMessage)
		if err != nil {
			return err
		}
		c, err2 := aes.NewCipher([]byte(FeiShu.EncryptKey))
		if err2 != nil {
			return err2
		}
		logrus.Info("密文：", encodeMessage.Encrypt)
		baseDecode, err3 := base64.StdEncoding.DecodeString(encodeMessage.Encrypt)
		if err3 != nil {
			return err3
		}
		c.Decrypt(body, baseDecode)
	}

	err = json.Unmarshal(body, &event)
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
	err = json.Unmarshal(body, &buttonData)
	if err != nil {
		return
	}
	buttonCallBack(buttonData)
	return
}
