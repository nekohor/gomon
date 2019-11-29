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

func RunServer(app *Application) {

	if app.Ctx.Cfg.ServerMode == "tcp" {
		RunTcpServer(app)
	} else if app.Ctx.Cfg.ServerMode == "http" {
		RunHttpServer(app)
	} else {
		log.Fatal("invalid server mode")
	}
}

//func RespondData(request []byte) map[string]*Coil {
//	app := New()
//
//	config := NewConfig()
//	config.SetComponentsDir("")
//	config.SetDataMaxNum(3999)
//	config.SetFmTagConfig(false)
//
//	app.Ctx.SetConfig(config)
//
//	coils := app.RespondCoils(string(request))
//
//	return coils
//}
