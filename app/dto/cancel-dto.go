package dto

type Cancel struct {
	Amt            string `json:"amt"`
	CancelMsg      string `json:"cancelMsg"`
	CancelRetryCnt string `json:"cancelRetryCnt"`
	CancelServerIP string `json:"cancelServerIp"`
	CancelType     string `json:"cancelType"`
	CancelUserID   string `json:"cancelUserId"`
	CancelUserInfo string `json:"cancelUserInfo"`
	CancelUserIP   string `json:"cancelUserIp"`
	IMid           string `json:"iMid"`
	MerchantToken  string `json:"merchantToken"`
	PayMethod      string `json:"payMethod"`
	PreauthToken   string `json:"preauthToken"`
	TXid           string `json:"tXid"`
	TimeStamp      string `json:"timeStamp"`
	Worker         string `json:"worker"`
	ReferenceNo    string `json:"referenceNo"`
}
