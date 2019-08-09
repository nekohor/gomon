package gomon

import (
	"log"
	"os"
	"github.com/BurntSushi/toml"
	"fmt"
	"time"
)

type Setting struct {
	Data struct {
		DataMaxNum int `toml:"DataMaxNum"`
	} `toml:"Data"`
	Mode struct {
		AppMode string `toml:"AppMode"`
	} `toml:"Mode"`
	TCPServer struct {
		Port int `toml:"port"`
	} `toml:"TcpServer"`
	BatchExportMode struct {
		IsFactorsSpecific bool   `toml:"IsFactorsSpecific"`
		Line              int    `toml:"Line"`
		RootDir1580       string `toml:"RootDir1580"`
		RootDir2250       string `toml:"RootDir2250"`
		ResultDir         string `toml:"ResultDir"`
		StartDate         int    `toml:"StartDate"`
		EndDate           int    `toml:"EndDate"`
	} `toml:"BatchExportMode"`
}


func NewSetting() *Setting {
	s := new(Setting)
	confPath := fmt.Sprintf(GetComponentsDir() + "/Setting.toml")
	if _, err := toml.DecodeFile(confPath, &s); err != nil {
		CheckError(err)
	}
	log.Println(s)

	return s
}

func (s *Setting) IsBatchExportMode() bool {
	return s.Mode.AppMode == "BatchExport"
}

func (s *Setting) GetRootDir() string {
	if s.BatchExportMode.Line == 1580 {
		return s.BatchExportMode.RootDir1580
	} else if s.BatchExportMode.Line == 2250 {
		return s.BatchExportMode.RootDir2250
	} else {
		panic("Line of BatchExportMode in setting for root dir is wrong")
	}
}

func (s *Setting) GetResultDir() string {
	if s.IsBatchExportMode() {
		return s.BatchExportMode.ResultDir
	} else {
		return os.Args[2]
	}
}

func (s *Setting) GetDateArray() []string {

	const layout = "20060102"
	start, err := time.Parse(layout, string(s.BatchExportMode.StartDate))
	if err != nil {
		log.Fatalf("Parse StartDate error: %v", err)
	}
	end, err := time.Parse(layout, string(s.BatchExportMode.EndDate))
	if err != nil {
		log.Fatalf("Parse EndDate error: %v", err)
	}

	duration, _ := time.ParseDuration("24h")
	sumDay := int(end.Sub(start).Hours()/24) + 1
	// log.Println(sumDay.Hours() / 24)

	dateArray := []string{}
	curDate, _ := time.Parse(layout, string(s.BatchExportMode.StartDate))
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

func (s *Setting) GetResultFilePathInBatchMode(curDate string) string {

	fileName := fmt.Sprintf("ExportedData_%s_%s.json", s.BatchExportMode.Line, curDate)
	fileDir := fmt.Sprintf("%s/%s/%s", s.GetResultDir(), s.BatchExportMode.Line, curDate[:6])
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
