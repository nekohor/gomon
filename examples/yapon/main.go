package main

import (
	"flag"
	"github.com/nekohor/gomon"
)


var (
	componentsDir string
	dataMaxNum int

	// server mode
	serverMode string
	port int

	//export mode
	export bool
	exportFrom string
	exportDest string

	isUseFmPond bool
)


func init() {
	flag.StringVar(&componentsDir,"c","","`components directory` or config")
	flag.IntVar(&dataMaxNum, "n", 3999, "max `number` of length data")
	flag.BoolVar(&isUseFmPond,"fm",false,"pond data used from `fm or fx` dca file")

	flag.StringVar(&serverMode,"serve","http","`serve` in http protocol(tcp or http)")
	flag.IntVar(&port,"p",8999,"`port` number")

	flag.BoolVar(&export, "export", false, "`export`")
	flag.StringVar(&exportFrom,"from","","`export from` a directory with coils")
	flag.StringVar(&exportDest,"dest","","`export destination` to a directory placing json file")


}

func main() {
	flag.Parse()

	config := gomon.NewConfig()

	config.SetComponentsDir(componentsDir)
	config.SetDataMaxNum(dataMaxNum)
	config.SetFmTagConfig(isUseFmPond)

	config.SetServerMode(serverMode)
	config.SetPort(port)

	if export {
		config.SetExportConfig(export, exportFrom, exportDest)
	}

	gomon.Run(config)
}
