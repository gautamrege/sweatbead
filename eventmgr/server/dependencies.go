package server

import (
	"github.com/gautamrege/sweatbead/eventmgr/app"
	"github.com/gautamrege/sweatbead/eventmgr/db"
	"github.com/gautamrege/sweatbead/eventmgr/sweat"
)

type dependencies struct {
	SweatService sweat.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
  logger := app.GetLogger()
  dbStore := db.NewStorer(appDB)

	sweatService := sweat.NewService(dbStore, logger)

	return dependencies{
		SweatService: sweatService,
	}, nil
}
