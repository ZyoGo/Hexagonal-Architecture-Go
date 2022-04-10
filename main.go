package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	go func() {
		adddress := fmt.Sprintf(":%d", 4001)
		if err := e.Start(adddress); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
