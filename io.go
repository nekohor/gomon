package gomon

import (
    "io/ioutil"
    "encoding/json"
    "os"
    "log"
    // "bufio"
)

func SaveFile(filename string, saveData []byte) bool {
    err := ioutil.WriteFile(filename, saveData, os.ModeAppend)
    if err != nil {
        return false
    }
    return true
}

func SaveJson(coils map[string]*Coil, filePath string) bool {
    b, err := json.MarshalIndent(coils, "", "  ")
    if err != nil {
        log.Fatal("json err:", err)
    }
    return SaveFile(filePath, b)
}

// func GetInput(input string) string {
//     reader := bufio.NewReader(input)
//     command, _, _ := reader.ReadLine()
//     return string(command)
// }