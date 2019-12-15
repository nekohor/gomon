package gomon

import "log"

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

func (app *Application) Stat(req *StatsRequest) *Coil {

	app.Ctx.Current.CurCoilId = req.CoilInfo.CoilId
	app.Ctx.Current.CurDir = req.CoilInfo.CurDir
	app.Ctx.Current.CurFactorName = req.CoilInfo.FactorName
	f := NewFactor(app.Ctx)




	return coil
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
