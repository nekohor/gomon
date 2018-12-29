package gomon

import (
    // "os"
    // "path/filepath"
    // "log"
)

type Config struct {
    Setting *Setting
    PartTable *PartTable
    FactorTable *FactorTable
    DLLCaller *DLLCaller

    CurCoilId string
    CurFactorName string
    // curPartNameOS string
    // curPartNameCL string
    // curPartNameDS string
}

func NewConfig() *Config {
    cfg := new(Config)

    cfg.Setting = NewSetting()
    cfg.PartTable = NewPartTable(cfg.Setting.ExeDir)
    cfg.FactorTable = NewFactorTable(cfg.Setting.ExeDir)
    cfg.DLLCaller = &DLLCaller{cfg.Setting.ExeDir + "/ReadDCADLL.dll"}

    return cfg
}

func (this *Config) ConcatPath(dcaFileName string) string {
    return this.Setting.CurDir + "/" + this.CurCoilId + "/" + dcaFileName + ".dca"
}