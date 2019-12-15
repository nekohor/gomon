package gomon

import "log"

func RunServer(app *Application) {

	if app.Ctx.Cfg.ServerMode == "tcp" {
		RunTcpServer(app)
	} else if app.Ctx.Cfg.ServerMode == "http" {
		RunHttpServer(app)
	} else {
		log.Fatal("invalid server mode")
	}
}