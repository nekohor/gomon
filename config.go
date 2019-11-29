package gomon

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Config struct {
	ComponentsDir string
	DataMaxNum    int

	ServerMode string
	Port       int

	//export
	IsExport   bool
	ExportFrom string
	ExportDest string

	FmTag string
}

func NewConfig() *Config {
	config := new(Config)
	return config
}

func (c *Config) SetComponentsDir(dir string) {
	c.ComponentsDir = dir
}

func (c *Config) GetComponentsDir() string {
	if c.ComponentsDir == "" {

	} else {
		if !strings.HasSuffix(c.ComponentsDir, "/") {
			c.ComponentsDir += "/"
		}
	}
	return GetExeDir() + "components"
}

func (c *Config) SetDataMaxNum(num int) {
	c.DataMaxNum = num
}

func (c *Config) SetFmTagConfig(isUse bool) {
	if isUse {
		c.FmTag = "_fm"
	} else {
		c.FmTag = ""
	}
}

func (c *Config) SetServerMode(mode string) {
	if mode != "tcp" && mode != "http" {
		log.Fatal("invalid server mode")
	} else {
		c.ServerMode = mode
	}
}

func (c *Config) SetPort(port int) {
	c.Port = port
}

func (c *Config) GetPort() string {
	return ":" + strconv.Itoa(c.Port)
}

func (c *Config) SetExportConfig(export bool, from string, dest string) {
	c.IsExport = export
	c.ExportFrom = from
	c.ExportDest = dest
}

// getter
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
