package gomon

import (
	"log"
)

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


func RespondData(request []byte) map[string]*Coil {
	app := New()

	config := NewConfig()
	config.SetComponentsDir("./")
	config.SetDataMaxNum(3999)

	app.Ctx.SetConfig(config)

	coils := app.RespondCoils(string(request))

	return coils
}


