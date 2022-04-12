package dto

type Payment struct {
	Amt           string `json:"amt"`
	CallBackURL   string `json:"callBackUrl"`
	ReferenceNo   string `json:"referenceNo"`
	TXid          string `json:"tXid"`
	MerchantToken string `json:"merTok"`
	TimeStamp     string `json:"timeStamp"`
}
