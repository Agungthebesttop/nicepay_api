package main

import (
	"fmt"
	"nicepay-api/app/config"
	"nicepay-api/app/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	config.LoadEnv()
}

func main() {
	e := echo.New()
	routes.NewRoutes(e)

	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
