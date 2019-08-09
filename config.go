package gomon

import (
	"fmt"
	"log"
)

type Config struct {
	ComponentsDir string
	Port int
	DataMaxNum int

	//export
	IsExport bool
	ExportFrom string
	ExportDest string
}


func NewConfig() *Config {
	config := new(Config)
	return config
}

func (c *Config) SetComponentsDir(dir string) {
	c.ComponentsDir = dir
}

func (c *Config) GetComponentsDir() string {
	return c.ComponentsDir + "/" + "components"
}

func (c *Config) SetPort(port int) {
	c.Port = port
}

func (c *Config) SetDataMaxNum(num int) {
	c.DataMaxNum = num
}

func (c *Config) SetExportConfig(export bool, from string, dest string) {
	c.IsExport = export
	c.ExportFrom = from
	c.ExportFrom = dest
}

func (c *Config) GetResultDir() string {
	return c.ExportDest
}

func (c *Config) GetResultFilePath() string {
	fileName := fmt.Sprintf("ExportedData.json")
	fileDir := c.GetResultDir()
	err := CreateDir(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	filePath := fileDir + "/" + fileName
	return filePath
}