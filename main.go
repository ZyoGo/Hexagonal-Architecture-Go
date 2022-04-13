package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	api "github.com/w33h/Hexagonal-Architecture-Go/api"
	modules "github.com/w33h/Hexagonal-Architecture-Go/app/modules"
	config "github.com/w33h/Hexagonal-Architecture-Go/config"
	util "github.com/w33h/Hexagonal-Architecture-Go/util"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisterModules(dbCon)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		adddress := fmt.Sprintf("localhost:%d", 4001)
		if err := e.Start(adddress); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
