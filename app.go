package gomon

import (
	//"fmt"
	"log"
	cache "github.com/patrickmn/go-cache"

)

type Application struct {
	Ctx *Context
}

func New() *Application {
	app := new(Application)
	app.Ctx = NewContext()
	return app
}

func (app *Application) RespondCoil(req *coilRequest) *Coil {

	app.Ctx.Current.CurCoilId = req.CoilId
	app.Ctx.Current.CurDir = req.CurDir
	app.Ctx.Current.CurFactorNames = req.FatcorNames
	coil := NewCoil(app.Ctx)

	return coil
}

func (app *Application) Stat(req *StatsRequest) string {

	app.Ctx.Current.CurCoilId = req.CoilInfo.CoilId
	app.Ctx.Current.CurDir = req.CoilInfo.CurDir
	app.Ctx.Current.CurFactorName = req.CoilInfo.FactorName

	key := req.CoilInfo.CoilId + "/" + req.CoilInfo.FactorName

	var s *Stats
	data, found := app.Ctx.CachePool.Get(key)
	if found {
		s = NewStats(app.Ctx, data.([]DataType), req)
	} else {
		f := NewFactor(app.Ctx)
		app.Ctx.CachePool.Set(key, f.Data, cache.DefaultExpiration)
		s =  NewStats(app.Ctx, f.Data, req)
	}

	return s.Calculate()
}

func (app *Application) ExportCurrent() {
	var coils map[string]*Coil

	curCoilIds := WalkDir(app.Ctx.Current.CurDir)
	coils = app.ExportCoils(curCoilIds)

	saveFilePath := app.Ctx.Cfg.GetResultFilePath()
	SaveJson(coils, saveFilePath)

	log.Println(saveFilePath)
}

//GetCoils 在web api中只用这一种，包括单卷的情况
func (app *Application) ExportCoils(resCoilIds []string) map[string]*Coil {

	coils := make(map[string]*Coil)

	for _, coilId := range resCoilIds {

		coils[coilId] = app.ExportCoil(coilId)

	}
	return coils
}

func (app *Application) ExportCoil(coilId string) *Coil {

	app.Ctx.Current.CurCoilId = coilId
	app.Ctx.Current.CurDir = app.Ctx.Cfg.ExportFrom
	app.Ctx.Current.CurFactorNames = app.Ctx.FactorConf.GetFactorNames()

	coil := NewCoil(app.Ctx)
	return coil
}

//func (app *Application) ExportBatch() {
//	var coils map[string]*Coil
//	for _, curDate := range app.Ctx.Config.GetDateArray() {
//
//		app.Ctx.Current.CurDir= app.Ctx.Config.GetCurDirInBatchMode(curDate)
//		curCoilIds := WalkDir(app.Ctx.Current.CurDir)
//
//		coils = app.GetCoils(curCoilIds)
//		saveFilePath := app.Ctx.Config.GetResultFilePathInBatchMode(curDate)
//		log.Println(saveFilePath)
//		SaveJson(coils, saveFilePath)
//	}
//}
