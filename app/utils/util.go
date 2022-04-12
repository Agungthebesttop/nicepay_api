package utils

import (
	"crypto/sha256"
	"fmt"
	"net"
	"os"
)

func GetIP() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)

	var ip string

	for _, addr := range addrs {
		ip = addr.To4().String()
	}

	return ip
}

func GenerateMerchantToken(timestampTrx, referenceNo, amt string) string {
	imId := os.Getenv("IMID")
	merchantKey := os.Getenv("MERCHANT_KEY")

	data := []byte(fmt.Sprintf("%s%s%s%s%s", timestampTrx, imId, referenceNo, amt, merchantKey))

	merchantToken := sha256.Sum256(data)

	return fmt.Sprintf("%x", merchantToken[:])
}
