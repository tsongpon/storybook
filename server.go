package main

import (
	"github.com/tsongpon/yoneebook/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/ping", handler.Ping)
	e.GET("/stories", handler.GetStory)

	e.Logger.Fatal(e.Start(":5000"))
}
