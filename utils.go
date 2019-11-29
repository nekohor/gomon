package gomon

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func WalkDir(path string) []string {
	fmt.Println(path)
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		log.Println("theDirPath cannot walk in utils.go:57")
		return []string{}
	}
	coilIdList := []string{}
	for _, fi := range rd {
		if fi.IsDir() {
			coilIdList = append(coilIdList, fi.Name())
		}
	}
	return coilIdList
}

func If(condition bool, trueVal, falseVal []string) []string {
	if condition {
		return trueVal
	}
	return falseVal
}

func GetMillLine(ctx *Context, coilId string) string {
	if string(coilId[0]) == "M" {
		return "1580"
	} else if string(coilId[0]) == "H" {
		return "2250"
	} else if string(coilId[0]) == "G" {
		return "2250"
	} else if string(coilId[0]) == "C" {
		return "1580"
	} else {
		log.Println("In JudgeLine Else Logic")
		log.Println(coilId)
		panic("This coil from wrong line.")
	}
}

//调用os.MkdirAll递归创建文件夹
func CreateDir(path string) error {
	if !IsExist(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

/* 判断所给路径文件/文件夹是否存在(返回true是存在) */
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetExeDir() string {
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return exeDir
}

func GetAbsPath(path string) string {
	absDir, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	return absDir
}

func GetComponentsDir() string {
	return GetExeDir() + "/" + "components"
}

func GJsonArrayToStringArray(gJsonArr []gjson.Result) []string {
	strArr := make([]string, len(gJsonArr))
	for i, v := range gJsonArr {
		strArr[i] = v.String()
	}
	return strArr
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func GetFmStands() []int {
	return []int {1,2,3,4,5,6,7}
}