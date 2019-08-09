package gomon

import (
// "os"
// "path/filepath"
// "log"
)

type Context struct {
	Cfg     *Config
	PartConf   *PartConfig
	FactorConf *FactorConfig
	Reader     *Reader

	CoilIds []string
	CurDir string

	//CurCoilId     string
	//CurFactorName string
}

func NewContext() *Context {
	this := new(Context)
	this.Reader = NewReader()
	return this
}

func (c * Context) SetConfig(config *Config) {
	c.Cfg = config
	c.PartConf = NewPartConfig(c.Cfg)
	c.FactorConf = NewFactorConfig(c.Cfg)
}
