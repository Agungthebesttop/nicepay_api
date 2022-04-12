package nicepay

import (
	"encoding/json"
	"fmt"
	"nicepay-api/app/entity"
	"os"
	"time"

	"github.com/vicanso/go-axios"
)

func doCheckStatus(dto []byte) (int, entity.Status) {
	baseURL := os.Getenv("STAGING_API_URL")

	if os.Getenv("APP_MODE") == "staging" {
		baseURL = os.Getenv("DEV_API_URL")
	}

	axi := axios.NewInstance(&axios.InstanceConfig{
		BaseURL: baseURL,
		Timeout: 30 * time.Second,
	})

	resp, err := axi.Post("/nicepay/direct/v2/inquiry", dto)
	if err != nil {
		fmt.Println(err)
	}

	var respBody entity.Status

	json.Unmarshal(resp.Data, &respBody)

	return resp.Status, respBody
}
