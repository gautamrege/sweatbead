package main

import (
	"fmt"
	"strconv"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/service"

	"github.com/urfave/negroni"
)

func main() {
	// mux router
	router := service.InitRouter()

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := 33001 // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	// Testing DB operation
	s := db.Sweat{
		Glucose:         1.3,
		Chloride:        2.1,
		Sodium:          1.2,
		HeartBeat:       72,
		RoomTemperature: 76,
	}

	_ = s.Create()

	server.Run(addr)

}
