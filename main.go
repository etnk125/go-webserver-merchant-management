package main

import (
	"github.com/etnk125/go-webserver-merchant-management/route"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	route.Init(e)
	e.Logger.Fatal(e.Start(":8080"))

}
