package main

import (
	"flag"
	"github.com/nekohor/gomon"
)


var (
	componentsDir string
	port int
	dataMaxNum int

	//export mode
	export bool
	exportFrom string
	exportDest string
)


func init() {
	flag.StringVar(&componentsDir,"c","./","`components directory` or config")
	flag.IntVar(&port,"s",8999,"`serve` with port number")
	flag.IntVar(&dataMaxNum, "n", 3999, "`max number` of length data")

	flag.BoolVar(&export, "export", false, "`export`")
	flag.StringVar(&exportFrom,"from","","`export from` a directory with coils")
	flag.StringVar(&exportDest,"dest","","`export destination` to a directory placing json file")

}

func main() {
	flag.Parse()

	config := gomon.NewConfig()

	config.SetComponentsDir(componentsDir)
	config.SetPort(port)
	config.SetDataMaxNum(dataMaxNum)

	if export {
		config.SetExportConfig(export, exportFrom, exportDest)
	}


	gomon.Run(config)
}
