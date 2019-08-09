package gomon

import (
	// "os"
	// "path/filepath"
	"log"
)

type Monitor struct {
	Ctx *Context
}

func New() *Monitor {
	g := new(Monitor)
	g.Ctx = NewContext()
	return g
}

func (g *Monitor) RespondCoils(req string) map[string]*Coil {
	sockConf := NewSocketConfig(req)

	coils := make(map[string]*Coil)

	for _, coilId := range sockConf.GetCoilIds() {
		g.Ctx.CurDir = sockConf.GetCurDir(coilId)
		factorNames := sockConf.GetFactors(coilId)
		log.Print(g.Ctx.CurDir)
		log.Print(coilId)
		log.Print(factorNames)
		coils[coilId] = NewCoil(g.Ctx, coilId, factorNames)
	}

	return coils
}

func (g *Monitor) ExportCurrent() {
	var coils map[string]*Coil

	g.Ctx.CurDir = g.Ctx.Cfg.ExportFrom
	curCoilIds := WalkDir(g.Ctx.CurDir)

	coils = g.GetCoils(curCoilIds)
	saveFilePath := g.Ctx.Cfg.GetResultFilePath()
	log.Println(saveFilePath)
	SaveJson(coils, saveFilePath)
}

func (g *Monitor) GetCoil(coilId string) *Coil {
	factorNames := g.Ctx.FactorConf.GetFactorNames()
	coil := NewCoil(g.Ctx, coilId, factorNames)
	return coil
}

//GetCoils 在web api中只用这一种，包括单卷的情况
func (g *Monitor) GetCoils(resCoilIds []string) map[string]*Coil {
	coils := make(map[string]*Coil)
	for _, coilId := range resCoilIds {
		coils[coilId] = g.GetCoil(coilId)
	}
	return coils
}

//func (g *Monitor) ExportBatch() {
//	var coils map[string]*Coil
//	for _, curDate := range g.Ctx.Config.GetDateArray() {
//
//		g.Ctx.CurDir = g.Ctx.Config.GetCurDirInBatchMode(curDate)
//		curCoilIds := WalkDir(g.Ctx.CurDir)
//
//		coils = g.GetCoils(curCoilIds)
//		saveFilePath := g.Ctx.Config.GetResultFilePathInBatchMode(curDate)
//		log.Println(saveFilePath)
//		SaveJson(coils, saveFilePath)
//	}
//}


