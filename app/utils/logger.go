package utils

import (
	"fmt"
	"os"
)

func LogReq(newLine interface{}) {
	logger("request", newLine)
}

func LogRes(newLine interface{}) {
	logger("response", newLine)
}

func logger(log_type, newLine interface{}) {
	f, err := os.OpenFile(fmt.Sprintf("%s.log", log_type), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
	}
}
