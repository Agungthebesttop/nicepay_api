package dto

type Registration struct {
	Amt             string      `json:"amt"`
	BankCd          string      `json:"bankCd"`
	BillingAddr     string      `json:"billingAddr"`
	BillingCity     string      `json:"billingCity"`
	BillingCountry  string      `json:"billingCountry"`
	BillingEmail    string      `json:"billingEmail"`
	BillingNm       string      `json:"billingNm"`
	BillingPhone    string      `json:"billingPhone"`
	BillingPostCd   string      `json:"billingPostCd"`
	BillingState    string      `json:"billingState"`
	CartData        interface{} `json:"cartData"`
	Currency        string      `json:"currency"`
	DBProcessURL    string      `json:"dbProcessUrl"`
	DeliveryAddr    string      `json:"deliveryAddr"`
	DeliveryCity    string      `json:"deliveryCity"`
	DeliveryCountry string      `json:"deliveryCountry"`
	DeliveryNm      string      `json:"deliveryNm"`
	DeliveryPhone   string      `json:"deliveryPhone"`
	DeliveryPostCd  string      `json:"deliveryPostCd"`
	DeliveryState   string      `json:"deliveryState"`
	Description     string      `json:"description"`
	Fee             string      `json:"fee"`
	GoodsNm         string      `json:"goodsNm"`
	IMid            string      `json:"iMid"`
	InstmntMon      string      `json:"instmntMon"`
	InstmntType     string      `json:"instmntType"`
	MerFixAcctID    string      `json:"merFixAcctId"`
	MerchantToken   string      `json:"merchantToken"`
	MitraCd         string      `json:"mitraCd"`
	NotaxAmt        string      `json:"notaxAmt"`
	PayMethod       string      `json:"payMethod"`
	RecurrOpt       string      `json:"recurrOpt"`
	ReferenceNo     string      `json:"referenceNo"`
	ReqClientVer    string      `json:"reqClientVer"`
	ReqDomain       string      `json:"reqDomain"`
	ReqDt           string      `json:"reqDt"`
	ReqServerIP     string      `json:"reqServerIP"`
	ReqTm           string      `json:"reqTm"`
	TimeStamp       string      `json:"timeStamp"`
	UserAgent       string      `json:"userAgent"`
	UserIP          string      `json:"userIP"`
	UserLanguage    string      `json:"userLanguage"`
	UserSessionID   string      `json:"userSessionID"`
	VacctValidDt    string      `json:"vacctValidDt"`
	VacctValidTm    string      `json:"vacctValidTm"`
	Vat             string      `json:"vat"`
}
