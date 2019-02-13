package gomon


import (
    "os"
    "path/filepath"
    "log"
    // "gopkg.in/yaml.v2"
    "github.com/go-ini/ini"
    // "io/ioutil"
    "time"
    "fmt"
    "strings"
)


type Setting struct {
    BatchMode bool `yaml:"batch_mode"`
    SpecificFactorsMode bool `yaml:"specific_factors_mode"`
    Factors []string `yaml:"factors"`

    TablesDir string `yaml:"tables_dir"`
    ExeDir string `yaml:"exe_dir"`

    // Line 仅仅为RootDir服务，卷号的产线判断用JudgeLine函数
    Line string `yaml:"line"`
    RootDir1 string `yaml:"root_dir1"`
    RootDir2 string `yaml:"root_dir2"`
    RootDir string `yaml:"root_dir"`

    CurDir string `yaml:"cur_dir"`
   
    ResultDirDefault string `yaml:"result_dir_default"`
    ResultDir string `yaml:"result_dir"`

    StartDate string `yaml:"start_date"`
    EndDate string `yaml:"end_date"`
    DateArray []string `yaml:"-"`

    CoilIds []string `yaml:"-"`

    MaxArray int `yaml:"max_array"`
}


func NewSetting() *Setting {
    s := new(Setting)
    s.ExeDir = GetExeDir()
    setup, err := ini.Load(s.ExeDir + "/setup.ini")
    if err != nil {
        log.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

    // data module
    s.MaxArray,err = setup.Section("data").Key("max_array").Int()
    if err != nil {
        log.Println("max_array err", err)
    }

    // tables module
    s.TablesDir = setup.Section("tables").Key("tables_dir").String()

    // specific module
    s.SpecificFactorsMode, err = setup.Section("specific").Key("specific_factors_mode").Bool()
    if err != nil {
        log.Println("mode err", err)
    }
    s.Factors = strings.Split(setup.Section("specific").Key("factors").String(), ",")

    // batch module
    s.BatchMode, err = setup.Section("batch").Key("batch_mode").Bool()
    if err != nil {
        log.Println("mode err", err)
    }

    s.Line = setup.Section("batch").Key("line").String()
    s.RootDir1 = setup.Section("batch").Key("root_dir1").String()
    s.RootDir2 = setup.Section("batch").Key("root_dir2").String()
    s.RootDir = s.GetRootDir()

    s.ResultDirDefault = setup.Section("batch").Key("result_dir").String()
    s.ResultDir = s.GetResultDir()

    s.StartDate = setup.Section("date").Key("start_date").String()
    s.EndDate = setup.Section("date").Key("end_date").String()
    s.DateArray = s.GetDateArray()
    log.Println(s.DateArray)

    s.CurDir = s.GetCurDir()
    log.Println(s.CurDir)
    s.CoilIds = s.GetCoilIdsInCurDir()

    return s
}

func GetExeDir() string {
    exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    return exeDir
}

func (s *Setting) GetRootDir() string {
    if s.Line == "1580" {
        return s.RootDir1
    } else if s.Line == "2250" {
        return s.RootDir2
    } else {
        panic("setup line is wrong")
    }
}

func (s *Setting) GetResultDir() string {
    if s.BatchMode {
        return s.ResultDirDefault
    } else {
        return s.GetResultDirFromOsArgs()
    }
}

func (s *Setting) GetResultDirFromOsArgs() string {
    ResultDir, err := filepath.Abs(filepath.Dir(os.Args[2]))
    if err != nil {
        log.Fatal(err)
    }
    return ResultDir
}

func (s *Setting) GetCurDir() string {
    if s.BatchMode {
        return s.GetCurDirFromDate(s.DateArray[0])
    } else {
        return s.GetCurDirFromOsArgs()
    }
}

func (s *Setting) GetCurDirFromDate(date string) string {
    return fmt.Sprintf("%s/%s/%s",s.RootDir,date[:6],date)
}

func (s *Setting) GetCurDirFromOsArgs() string {
    curDir, err := filepath.Abs(filepath.Dir(os.Args[1]))
    if err != nil {
        log.Fatal(err)
    }
    return curDir
}

func (s *Setting) GetCoilIdsInCurDir() []string {
    return WalkDir(s.CurDir)
}

func (s *Setting) GetDateArray() []string {
    const layout = "20060102"
    start, err := time.Parse(layout, s.StartDate)
    if err != nil {
        log.Fatalf("Parse StartDate error: %v", err)
    }
    end, err := time.Parse(layout, s.EndDate)
    if err != nil {
        log.Fatalf("Parse EndDate error: %v", err)
    }

    duration, _ := time.ParseDuration("24h")
    sumDay := int(end.Sub(start).Hours() / 24) + 1
    // log.Println(sumDay.Hours() / 24)
    dateArray := []string{}
    curDate, _ := time.Parse(layout, s.StartDate)
    for i := 0; i < sumDay; i++ {
        // log.Println("iter", dateArray)
        dateArray = append(dateArray, curDate.Format("20060102"))
        curDate = curDate.Add(duration)
        // log.Println(curDate)
    }
    return dateArray
}

func (s *Setting) GetFactorArray(defaultFactors []string) []string {
    // If(s.SpecificFactorsMode, s.Factors, defaultFactors)
    if s.SpecificFactorsMode {
        return s.Factors
    } else {
        return defaultFactors
    }
}

func (s *Setting) GetMillLine(coilId string) string {
    if s.BatchMode {
        return s.Line
    } else {
        return JudgeLine(coilId)
    }
}

func (s *Setting) GetSaveFilePath(date string) string {
    fileName := fmt.Sprintf("ExportedData_%s_%s.json",s.Line, date)
    fileDir := fmt.Sprintf("%s/%s/%s",s.ResultDir,s.Line,date[:6])
    err := CreateDir(fileDir)
    if err != nil {
        log.Fatal(err)
    }
    filePath := fileDir + "/" + fileName
    return filePath
}

func (s *Setting) GetDefaultFilePath() string {
    fileName := fmt.Sprintf("ExportedData.json")
    fileDir := s.ResultDir
    err := CreateDir(fileDir)
    if err != nil {
        log.Fatal(err)
    }
    filePath := fileDir + "/" + fileName
    return filePath
}