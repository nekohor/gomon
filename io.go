package gomon

import (
    "io/ioutil"
    // "encoding/json"
    "os"

)

func SaveFile(filename string, saveData []byte) bool {
    err := ioutil.WriteFile(filename, saveData, os.ModeAppend)
    if err != nil {
        return false
    }
    return true
}