package gomon

import (
	"log"
)

func Run(config *Config) {

	log.Println(config)

	app := New()
	app.Ctx.SetConfig(config)

	if app.Ctx.Cfg.IsExport {
		app.ExportCurrent()
	} else {
		RunServer(app)
	}
}

