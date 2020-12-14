package main

import (
	"github.com/jordanparker6/go-api/src/api"
	"github.com/jordanparker6/go-api/src/core"
)

var db core.DatabaseI

func init() {
	db = core.NewDatabase(
		core.Config.Database.Dialect,
		core.Config.Database.Uri,
	)
	db.Connect()
}

func main() {
	api.Run(
		core.Config.Server.Host,
		core.Config.Server.Port,
	)
}
