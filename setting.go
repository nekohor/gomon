package gomon


import (
    // "os"
    // "path/filepath"
    "log"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "time"
    "fmt"
)


type Setting struct {
    RootDir string `yaml:"root_dir"`
    StartDate string `yaml:"start_date"`
    EndDate string `yaml:"end_date"`
    DateArray []string `yaml:"-"`
}


func NewSetting(exeDir string) *Setting {
    data, err := ioutil.ReadFile(exeDir + "setup.yaml")
    if err != nil {
        panic("setup.yaml read fail")
    }

    s := Setting{}
    err = yaml.Unmarshal([]byte(data), &s)
    log.Println(s)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return &s
}

func (s *Setting) GetCurDir(date string) string {
    return fmt.Sprintf("%s/%s/%s",s.RootDir,date[:4],date)
}

func (s *Setting) GetDateArray() string {
    const gobirthday = "20060102"
    start, err := time.Parse(gobirthday, s.StartDate)
    if err != nil {
        log.Fatalf("Parse StartDate error: %v", err)
    }
    end, err := time.Parse(gobirthday, s.EndDate)
    if err != nil {
        log.Fatalf("Parse EndDate error: %v", err)
    }

    duration, _ := time.ParseDuration("24h")
    sumDay := int(end.Sub(start).Hours() / 24) + 1
    // log.Println(sumDay.Hours() / 24)
    s.DateArray = make([]string, sumDay)
    curDate, _ := time.Parse(gobirthday, s.StartDate)
    for i := 0; i < sumDay; i++ {
        s.DateArray = append(s.DateArray, curDate.Format("20060102"))
        curDate = curDate.Add(duration)
    }
    log.Println(s.DateArray)
    
    return string(sumDay)
}


