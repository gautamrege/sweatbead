package main

// @APITitle SweatBead API
// @APIDescription SweatBead API for Microservices in Go and MongoDB!

import (
	"fmt"
	"strconv"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/config"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/service"

	"github.com/urfave/negroni"
)

func main() {
	config.Load()

	logger.Init()
	db.Init()

	// mux router
	router := service.InitRouter()

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)

}
