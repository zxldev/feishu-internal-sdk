package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/zxldev/feishu-internal-sdk/core/consts"
	"github.com/zxldev/feishu-internal-sdk/core/model"
	"io/ioutil"
	"net/http"
	"time"
)

var FeiShu FeiShuInternal

type FeiShuInternal struct {
	AppId     string  `json:"app_id"`
	AppSerect string  `json:"app_serect"`
	Tenant    *Tenant `json:"tenant"`
}

func (f *FeiShuInternal) Init(AppId, AppSerect string) {
	FeiShu = FeiShuInternal{AppId: AppId, AppSerect: AppSerect, Tenant: &Tenant{}}
}

type Tenant struct {
	TenantAccessToken string    `json:"tenant_access_token"`
	Expire            time.Time `json:"expire"`
}

func (t *Tenant) GetTenantAccessToken() string {
	if t.TenantAccessToken != "" && t.Expire.Unix() > time.Now().Unix() {
		return t.TenantAccessToken
	} else {
		token, err := t.PostAuth(consts.ApiTenantAccessTokenInternal, map[string]string{
			"app_id":     FeiShu.AppId,
			"app_secret": FeiShu.AppSerect,
		})
		if err != nil {
			log.Error("获取token失败:", err.Error())
			return ""
		} else {
			t.TenantAccessToken = "Bearer " + token.TenantAccessToken
			t.Expire = time.Now().Add(time.Second * time.Duration(token.Expire-60)) //提前60秒结束
			return t.TenantAccessToken
		}
	}
}

func (t Tenant) PostAuth(url string, data interface{}) (r *model.AppAccessTokenInternalResp, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", consts.GetConst(url), bytes.NewBuffer(b))
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	rb, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}
	r = &model.AppAccessTokenInternalResp{}
	err = json.Unmarshal(rb, &r)
	if err != nil {
		return
	}
	if r.TenantAccessToken == "" {
		return nil, errors.New("获取token失败")
	} else {
		return
	}
}

func (t Tenant) Post(url string, data interface{}) (rb []byte, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", consts.GetConst(url), bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", t.GetTenantAccessToken())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	return ioutil.ReadAll(resp.Body)
}
