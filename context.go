package gomon

import (
// "os"
// "path/filepath"
// "log"
)

type Context struct {
	Setting    *Setting
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

	this.Setting = NewSetting()
	this.PartConf = NewPartConfig(this.Setting)
	this.FactorConf = NewFactorConfig(this.Setting)
	this.Reader = NewReader()

	this.CurDir = this.Setting.InitCurDir()
	this.CoilIds = WalkDir(this.CurDir)
	return this
}
