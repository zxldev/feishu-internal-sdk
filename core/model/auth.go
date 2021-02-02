package model

type AppAccessTokenInternalResp struct {
	Common
	TenantAccessToken string `json:"tenant_access_token"`
	AppAccessToken    string `json:"app_access_token"`
	Expire            int64  `json:"expire"`
}
