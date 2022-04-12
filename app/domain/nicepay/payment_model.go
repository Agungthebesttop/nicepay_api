package nicepay

import (
	"fmt"
	"net/http"
	"nicepay-api/app/dto"
	"os"
	"time"

	"github.com/vicanso/go-axios"
)

func doPayment(dto dto.Payment) (int, string) {
	baseURL := os.Getenv("STAGING_API_URL")

	if os.Getenv("APP_MODE") == "staging" {
		baseURL = os.Getenv("DEV_API_URL")
	}

	header := make(http.Header)
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	axi := axios.NewInstance(&axios.InstanceConfig{
		BaseURL: baseURL,
		Timeout: 30 * time.Second,
		Headers: header,
	})

	resp, err := axi.Post(fmt.Sprintf("/nicepay/direct/v2/payment?timeStamp=%s&tXid=%s&merchantToken=%s&callBackUrl=%s", dto.TimeStamp, dto.TXid, dto.MerchantToken, dto.CallBackURL), nil)
	if err != nil {
		fmt.Println(err)
	}

	return resp.Status, string(resp.Data)
}
