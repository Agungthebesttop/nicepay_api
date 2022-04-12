package nicepay

import (
	"encoding/json"
	"nicepay-api/app/dto"
	"nicepay-api/app/utils"
	"os"

	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	statusDto := new(dto.Status)

	statusDto.IMid = os.Getenv("IMID")

	if err := c.Bind(statusDto); err != nil {
		return err
	}

	statusDto.MerchantToken = utils.GenerateMerchantToken(statusDto.TimeStamp, statusDto.ReferenceNo, statusDto.Amt)

	req, err := json.Marshal(statusDto)
	if err != nil {
		return err
	}

	// log request
	utils.LogReq(string(req))

	code, resp := doCheckStatus(req)

	jsonRes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// log response
	utils.LogRes(string(jsonRes))

	return c.JSON(code, resp)
}
