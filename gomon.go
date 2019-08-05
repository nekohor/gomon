package gomon

func Run() {
	app := NewMonitor()

	mode := app.Context.Setting.Mode.AppMode

	switch mode {
	case "DefaultExport":
		app.ExportDefault()
	case "BatchExport":
		app.ExportBatch()
	case "TcpServer":
		RunServer(app)
	}
}



