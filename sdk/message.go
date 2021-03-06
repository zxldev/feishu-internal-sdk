package sdk

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zxldev/feishu-internal-sdk/core/consts"
	"github.com/zxldev/feishu-internal-sdk/core/model"
	"github.com/zxldev/feishu-internal-sdk/core/util/encrypt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (t Tenant) SendMessage(msg model.Msg) {
	r, err := t.Post(consts.ApiRobotSendMessage, msg)
	if err != nil {
		logrus.Info(err.Error())
	} else {
		logrus.Info(string(r))
	}
}

func (t Tenant) MessageCardCallback(w http.ResponseWriter, r *http.Request, action func(button model.EventAction)) (err error) {
	event := model.Event{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	logrus.Info(string(body))

	err = json.Unmarshal(body, &event)
	if err != nil {
		return
	}
	//if event.Token != FeiShu.VerificationToken {
	//	w.WriteHeader(404)
	//	return
	//}
	if event.Type != "" {
		if event.Type == "url_verification" {
			w.Write([]byte(`{"challenge":"` + event.Challenge + `"}`))
			return
		}
	}

	//请求安全性校验
	//@see https://open.feishu.cn/document/ukTMukTMukTM/uYzMxEjL2MTMx4iNzETM
	timestamp := r.Header.Get("X-Lark-Request-Timestamp")
	nonce := r.Header.Get("X-Lark-Request-Nonce")

	var b strings.Builder
	b.WriteString(timestamp)
	b.WriteString(nonce)
	b.WriteString(FeiShu.VerificationToken)
	b.WriteString(string(body))
	bs := []byte(b.String())
	h := sha1.New()
	h.Write(bs)
	bs = h.Sum(nil)
	sig := fmt.Sprintf("%x", bs)

	if r.Header.Get("X-Lark-Signature") != sig {
		return errors.New("http.StatusUnauthorized")
	}

	//开始按照button 回调处理

	actionData := model.EventAction{}
	err = json.Unmarshal(body, &actionData)
	if err != nil {
		return
	}
	action(actionData)
	return
}

func (t Tenant) EventCallback(w http.ResponseWriter, r *http.Request, eventCallback func(event model.Event)) (err error) {
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

		body, err = encrypt.AesDecrypt(FeiShu.EncryptKey, encodeMessage.Encrypt)
		if err != nil {
			return err
		}
	}

	err = json.Unmarshal(body, &event)
	if err != nil {
		return
	}
	if event.Token != FeiShu.VerificationToken {
		w.WriteHeader(404)
		return
	}
	if event.Type != "" {
		if event.Type == "url_verification" {
			w.Write([]byte(`{"challenge":"` + event.Challenge + `"}`))
			return
		} else if event.Type == "event_callback" {
			eventCallback(event)
			return
		}
	}
	//TODO v2 的消息 https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM
	return
}
