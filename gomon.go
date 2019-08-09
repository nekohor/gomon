package gomon

import "log"

func Run(config *Config) {

	log.Println(config)

	app := New()
	app.Ctx.SetConfig(config)

	if config.IsExport {
		app.ExportCurrent()
	} else {
		RunServer(app)
	}
}



