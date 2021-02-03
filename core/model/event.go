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
