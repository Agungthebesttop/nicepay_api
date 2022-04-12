package nicepay

import (
	"encoding/json"
	"nicepay-api/app/dto"
	"nicepay-api/app/utils"

	"github.com/labstack/echo/v4"
)

func Payment(c echo.Context) error {
	paymentDto := new(dto.Payment)

	if err := c.Bind(paymentDto); err != nil {
		return err
	}

	// set token
	paymentDto.MerchantToken = utils.GenerateMerchantToken(
		paymentDto.TimeStamp,
		paymentDto.ReferenceNo,
		paymentDto.Amt)

	req, err := json.Marshal(paymentDto)
	if err != nil {
		return err
	}

	// log request
	utils.LogReq(string(req))

	code, resp := doPayment(*paymentDto)

	// log response
	utils.LogRes(resp)

	return c.HTML(code, resp)
}
