package gomon

import (
	"log"
	"os"
	//"path/filepath"
	// "gopkg.in/json.v2"
	"github.com/go-ini/ini"

	"fmt"
	"strings"
	"time"
)

type Setting struct {
	IniFile *ini.File

	DataMaxNum  int
	MaxPrintNum int

	SpecificFactorsMode bool
	SpecificFactors     []string

	Line string

	StartDate string
	EndDate   string
	DateArray []string
}

func NewSetting() *Setting {
	s := new(Setting)

	configuration, err := ini.Load(GetComponentsDir() + "/Configuration.ini")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	s.IniFile = configuration

	s.DataMaxNum, err = s.IniFile.Section("Data").Key("DataMaxNum").Int()
	if err != nil {
		log.Println("Parse DataMaxNum Error", err)
	}

	s.SpecificFactorsMode, err = s.IniFile.Section("Specific").Key("SpecificFactorsMode").Bool()
	if err != nil {
		log.Println("Parse SpecificFactorsMode Error", err)
	}
	s.SpecificFactors = strings.Split(s.IniFile.Section("Specific").Key("SpecificFactors").String(), ",")

	if s.IsBatchMode() {
		s.Line = s.IniFile.Section("Batch").Key("Line").String()
		s.StartDate = s.IniFile.Section("Batch").Key("StartDate").String()
		s.EndDate = s.IniFile.Section("Batch").Key("EndDate").String()
		s.DateArray = s.GetDateArray()
		log.Println(s.DateArray)
	}
	return s
}

func (s *Setting) IsBatchMode() bool {
	BatchMode, err := s.IniFile.Section("Batch").Key("BatchMode").Bool()
	if err != nil {
		log.Println("Parse BatchMode Error", err)
	}
	return BatchMode
}

func (s *Setting) GetRootDir() string {
	if s.Line == "1580" {
		return s.IniFile.Section("Batch").Key("RootDir1580").String()
	} else if s.Line == "2250" {
		return s.IniFile.Section("Batch").Key("RootDir2250").String()
	} else {
		panic("Setup line for root dir is wrong")
	}
}

func (s *Setting) GetResultDir() string {
	if s.IsBatchMode() {
		return s.IniFile.Section("Batch").Key("ResultDir").String()
	} else {
		return os.Args[2]
	}
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
	sumDay := int(end.Sub(start).Hours()/24) + 1
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

func (s *Setting) GetCurDirInBatchMode(date string) string {
	return fmt.Sprintf("%s/%s/%s", s.GetRootDir(), date[:6], date)
}

func (s *Setting) InitCurDir() string {
	if s.IsBatchMode() {
		return s.GetCurDirInBatchMode(s.StartDate)
	} else {
		return os.Args[1]
	}
}

func (s *Setting) GetResultFilePathInBatchMode(curDate string) string {

	fileName := fmt.Sprintf("ExportedData_%s_%s.json", s.Line, curDate)
	fileDir := fmt.Sprintf("%s/%s/%s", s.GetResultDir(), s.Line, curDate[:6])
	err := CreateDir(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	filePath := fileDir + "/" + fileName
	return filePath
}

func (s *Setting) GetResultFilePath() string {
	fileName := fmt.Sprintf("ExportedData.json")
	fileDir := s.GetResultDir()
	err := CreateDir(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	filePath := fileDir + "/" + fileName
	return filePath
}
