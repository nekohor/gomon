package gomon

import (
    "io/ioutil"
    "os"
    "log"
)

func WalkDir(path string) []string {
    rd, err := ioutil.ReadDir(path)
    if err != nil {
        log.Println("theDirPath cannot walk")
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

func JudgeLine(coilId string) string {
    if string(coilId[0]) == "M" {
        return "1580"
    } else if string(coilId[0]) == "H" {
        return "2250"
    } else {
        panic("This coil from wrong line.")
    }
}

//调用os.MkdirAll递归创建文件夹
func CreateDir(path string)  error  {
    if !IsExist(path) {
        err := os.MkdirAll(path,os.ModePerm)
        return err
    }
    return nil
}
 
// 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
    _, err := os.Stat(path)    //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}
