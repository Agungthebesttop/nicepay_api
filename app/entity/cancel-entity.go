package entity

type Cancel struct {
	Amt          string `json:"amt"`
	BankCd       string `json:"bankCd"`
	BillingNm    string `json:"billingNm"`
	Currency     string `json:"currency"`
	Description  string `json:"description"`
	GoodsNm      string `json:"goodsNm"`
	MitraCd      string `json:"mitraCd"`
	PayMethod    string `json:"payMethod"`
	PayNo        string `json:"payNo"`
	PayValidDt   string `json:"payValidDt"`
	PayValidTm   string `json:"payValidTm"`
	PaymentExpDt string `json:"paymentExpDt"`
	PaymentExpTm string `json:"paymentExpTm"`
	QrContent    string `json:"qrContent"`
	QrURL        string `json:"qrUrl"`
	ReferenceNo  string `json:"referenceNo"`
	RequestURL   string `json:"requestURL"`
	ResultCd     string `json:"resultCd"`
	ResultMsg    string `json:"resultMsg"`
	TXid         string `json:"tXid"`
	TransDt      string `json:"transDt"`
	TransTm      string `json:"transTm"`
	VacctNo      string `json:"vacctNo"`
	VacctValidDt string `json:"vacctValidDt"`
	VacctValidTm string `json:"vacctValidTm"`
}
