package model

type EncodeMessage struct {
	Encrypt string `json:"encrypt"`
}

type Event struct {
	Ts        float32 `json:"ts"`
	Uuid      string  `json:"uuid"`
	Token     string  `json:"token"`
	Type      string  `json:"type"`
	Challenge string  `json:"challenge"`
	Event     interface{}
}

type ButtonCallback struct {
	OpenId        string `json:"open_id"`
	UserId        string `json:"user_id"`
	OpenMessageId string `json:"open_message_id"`
	TenantKey     string `json:"tenant_key"`
	Token         string `json:"token"`
	Action        struct {
		Tag   string                 `json:"tag"`
		Value map[string]interface{} `json:"value"`
	} `json:"action"`
}
