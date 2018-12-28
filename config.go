package gomon

import (
    // "os"
    // "path/filepath"
    // "log"
)

type Config struct {
    PathConfig *PathConfig
    // CustomConfig *Setting
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

    cfg.PathConfig = NewPathConfig()
    // cfg.CustomConfig = NewSetting()

    cfg.PartTable = NewPartTable(cfg.PathConfig.ExeDir)
    cfg.FactorTable = NewFactorTable(cfg.PathConfig.ExeDir)
    cfg.DLLCaller = &DLLCaller{cfg.PathConfig.ExeDir + "/ReadDCADLL.dll"}

    return cfg
}

func (this *Config) ConcatPath(dcaFileName string) string {
    return this.PathConfig.CurDir + "/" + this.CurCoilId + "/" + dcaFileName + ".dca"
}