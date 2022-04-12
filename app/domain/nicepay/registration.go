package nicepay

import (
	"encoding/json"
	"nicepay-api/app/dto"
	"nicepay-api/app/utils"
	"os"

	"github.com/ibnusurkati/timeformat"
	"github.com/labstack/echo/v4"
)

func Registration(c echo.Context) error {
	timestampTrx := timeformat.Now("YYYYMMDDhmmss")
	referenceNum := generateReferenceNum(timestampTrx)

	registrationDto := new(dto.Registration)

	registrationDto.TimeStamp = timestampTrx
	registrationDto.IMid = os.Getenv("IMID")
	registrationDto.PayMethod = "02"
	registrationDto.ReferenceNo = referenceNum

	if err := c.Bind(registrationDto); err != nil {
		return err
	}

	// set token
	registrationDto.MerchantToken = utils.GenerateMerchantToken(
		timestampTrx,
		referenceNum,
		registrationDto.Amt)

	req, err := json.Marshal(registrationDto)
	if err != nil {
		return err
	}

	// log request
	utils.LogReq(string(req))

	code, resp := doRegistration(req)

	jsonRes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// log response
	utils.LogRes(string(jsonRes))

	return c.JSON(code, resp)
}
