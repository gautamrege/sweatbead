package server

import (
	"github.com/gautamrege/sweatbead/eventmgr/app"
	"github.com/gautamrege/sweatbead/eventmgr/db"
)

type dependencies struct {
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	_ = app.GetLogger()
	_ = db.NewStorer(appDB)

	return dependencies{}, nil
}
