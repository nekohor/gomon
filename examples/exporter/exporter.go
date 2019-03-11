package main

import (
	"github.com/nekohor/gomon"
)

func main() {
	app := gomon.NewGoMonitor()
	if app.Context.Setting.IsBatchMode() {
		app.ExportBatch()
	} else {
		app.ExportDefault()
	}
}
