package gomon

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "log"
)

type PathConfig struct {
    ExeDir string `json:"exeDir"`
    CurDir string `json:"curDir"`
    CoilIds []string `json:"coilIds"`
}

func NewPathConfig() *PathConfig {
    cfg := new(PathConfig)

    cfg.ExeDir = cfg.GetExeDir()
    cfg.CurDir = cfg.GetCurDir()

    log.Println(cfg.ExeDir)
    log.Println(cfg.CurDir)

    cfg.CoilIds = cfg.WalkDir(cfg.CurDir)
    return cfg
}

func (cfg *PathConfig) GetExeDir() string {
    exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    return exeDir
}

func (cfg *PathConfig) GetCurDir() string {
    curDir, err := filepath.Abs(filepath.Dir(os.Args[1]))
    if err != nil {
        log.Fatal(err)
    }
    return curDir
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
