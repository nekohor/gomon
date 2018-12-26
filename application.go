package gomon

import (
    // "os"
    // "path/filepath"
    // "log"
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

func (app *Application) GetCoils(resCoilIds []string) []Coil {
    coils := make(map[string]*Coil)
    var coil *Coil
    for _, coilId := range resCoilIds {
        coil = NewCoil(coilId)
        coil.PutData(app.Config)
        coils[coilId] = coil
    }
    return make([]Coil,len(resCoilIds))
}

