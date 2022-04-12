package nicepay

import (
	"encoding/json"
	"nicepay-api/app/dto"
	"nicepay-api/app/utils"
	"os"

	"github.com/labstack/echo/v4"
)

func Cancel(c echo.Context) error {
	cancelDto := new(dto.Cancel)

	cancelDto.IMid = os.Getenv("IMID")
	cancelDto.PayMethod = "02"

	if err := c.Bind(cancelDto); err != nil {
		return err
	}

	cancelDto.MerchantToken = utils.GenerateMerchantToken(cancelDto.TimeStamp, cancelDto.ReferenceNo, cancelDto.Amt)

	req, err := json.Marshal(cancelDto)
	if err != nil {
		return err
	}

	// log request
	utils.LogReq(string(req))

	code, resp := doCancelRegistration(req)

	jsonRes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// log response
	utils.LogRes(string(jsonRes))

	return c.JSON(code, resp)
}
