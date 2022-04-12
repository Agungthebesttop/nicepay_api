package dto

type Status struct {
	Amt           string `json:"amt"`
	IMid          string `json:"iMid"`
	MerchantToken string `json:"merchantToken"`
	ReferenceNo   string `json:"referenceNo"`
	TXid          string `json:"tXid"`
	TimeStamp     string `json:"timeStamp"`
}
