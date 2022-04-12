package nicepay

import (
	"encoding/json"
	"fmt"
	"nicepay-api/app/entity"
	"os"
	"time"

	"github.com/vicanso/go-axios"
)

func generateReferenceNum(timestampTrx string) string {
	return "AGUNG-" + timestampTrx
}

func doRegistration(dto []byte) (int, entity.Registration) {
	baseURL := os.Getenv("STAGING_API_URL")

	if os.Getenv("APP_MODE") == "staging" {
		baseURL = os.Getenv("DEV_API_URL")
	}

	axi := axios.NewInstance(&axios.InstanceConfig{
		BaseURL: baseURL,
		Timeout: 30 * time.Second,
	})

	resp, err := axi.Post("/nicepay/direct/v2/registration", dto)
	if err != nil {
		fmt.Println(err)
	}

	var respBody entity.Registration

	json.Unmarshal(resp.Data, &respBody)

	return resp.Status, respBody
}
