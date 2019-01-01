package main

import (
    "github.com/nekohor/gomon"
    "encoding/json"
    "log"
    // "fmt"
)


func main() {
    app := gomon.NewApp()
    coils := app.ExportAll()
    b, err := json.Marshal(coils)
    if err != nil {
        log.Println("json err:", err)
    }
    // log.Println(b)
    gomon.SaveFile(app.Config.Setting.ResultDir + "/reuslt.json", b)
}