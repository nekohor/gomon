package gomon

import (
    "log"
    "io/ioutil"
    "path/filepath"
)

type PathConfig struct {
    exeDir string `json:"-"`
    curDir string `json:"-"`
    coilIdList []string `json:"-"`
}

func NewPathConfig(osArgs0, osArgs1 string) *PathConfig {
    cfg := new(PathConfig)
    exeDir, err := filepath.Abs(filepath.Dir(osArgs0))
    if err != nil {
        panic(err)
    }
    cfg.exeDir = exeDir
    log.Println(cfg.exeDir)
    cfg.curDir = osArgs1
    cfg.coilIdList = cfg.WalkDir(cfg.curDir)
    return cfg
}

func (cfg *PathConfig) WalkDir(path string) []string {
    rd, err := ioutil.ReadDir(path)
    if err != nil {
        panic("theDirPath cannot walk")
    }
    coilIdList := []string{}
    for _, fi := range rd {
        if fi.IsDir() {
            coilIdList = append(coilIdList, fi.Name())
        }
    }
    return coilIdList
}