package gomon

import (
    // "os"
    // "path/filepath"
    "log"
)


type Application struct {
    Config *Config
}

func NewApp() *Application {
    app := new(Application)
    app.Config = NewConfig()
    return app
}

func (app *Application) GetCoil(coilId string) *Coil {
    coil := NewCoil(coilId)
    coil.PutData(app.Config)
    return coil
}

func (app *Application) GetCoils(resCoilIds []string) map[string]*Coil {
    coils := make(map[string]*Coil)
    for _, coilId := range resCoilIds {
        coils[coilId] = app.GetCoil(coilId)
    }
    return coils
}

func (app *Application) ExportDaily() {
    var coils map[string]*Coil
    for _, date := range app.Config.Setting.DateArray {

        app.Config.Setting.CurDir = app.Config.Setting.GetCurDirFromDate(date)
        curCoilIds := app.Config.Setting.GetCoilIdsInCurDir()

        coils = app.GetCoils(curCoilIds)
        // coils = make(map[string]*Coil)
        // for _, coilId := range curCoilIds {
        //     coils[coilId] = app.GetCoil(coilId)
        // }

        saveFilePath := app.Config.Setting.GetSaveFilePath(date)
        log.Println(saveFilePath)
        SaveJson(coils, saveFilePath)
    }
}

func (app *Application) ExportDefault() {
    var coils map[string]*Coil
    coils = app.GetCoils(app.Config.Setting.CoilIds)
    saveFilePath := app.Config.Setting.GetDefaultFilePath()
    log.Println(saveFilePath)
    SaveJson(coils, saveFilePath)
}
