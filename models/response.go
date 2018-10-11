package models

type Response struct {
	Code    int16       `json:"code"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message"`
	Reason  string      `json:"reason"`
}
