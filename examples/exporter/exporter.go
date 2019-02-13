package main

import (
    "github.com/nekohor/gomon"
)


func main() {
    app := gomon.NewApp()
    if app.Config.Setting.BatchMode {
        app.ExportDaily()
    } else {
        app.ExportDefault()
    }
}