package gomon

import (
	// "os"
	// "path/filepath"
	"log"
	"os"
)

type GoMonitor struct {
	Context *Context
}

func NewGoMonitor() *GoMonitor {
	g := new(GoMonitor)
	g.Context = NewContext()
	return g
}

func (g *GoMonitor) RespondCoils(req string) map[string]*Coil {
	sockConf := NewSocketConfig(req)

	coils := make(map[string]*Coil)

	for _, coilId := range sockConf.GetCoilIds() {
		g.Context.CurDir = sockConf.GetCurDir(coilId)
		factorNames := sockConf.GetFactors(coilId)
		log.Print(g.Context.CurDir)
		log.Print(coilId)
		log.Print(factorNames)
		coils[coilId] = NewCoil(g.Context, coilId, factorNames)
	}

	return coils
}

func (g *GoMonitor) GetCoil(coilId string) *Coil {
	factorNames := g.Context.FactorConf.GetFactorNames()
	coil := NewCoil(g.Context, coilId, factorNames)
	return coil
}

//GetCoils 在web api中只用这一种，包括单卷的情况
func (g *GoMonitor) GetCoils(resCoilIds []string) map[string]*Coil {
	coils := make(map[string]*Coil)
	for _, coilId := range resCoilIds {
		coils[coilId] = g.GetCoil(coilId)
	}
	return coils
}

func (g *GoMonitor) ExportBatch() {
	var coils map[string]*Coil
	for _, curDate := range g.Context.Setting.DateArray {

		g.Context.CurDir = g.Context.Setting.GetCurDirInBatchMode(curDate)
		curCoilIds := WalkDir(g.Context.CurDir)

		coils = g.GetCoils(curCoilIds)
		saveFilePath := g.Context.Setting.GetResultFilePathInBatchMode(curDate)
		log.Println(saveFilePath)
		SaveJson(coils, saveFilePath)
	}
}

func (g *GoMonitor) ExportDefault() {
	var coils map[string]*Coil

	g.Context.CurDir = os.Args[1]
	curCoilIds := WalkDir(g.Context.CurDir)

	coils = g.GetCoils(curCoilIds)
	saveFilePath := g.Context.Setting.GetResultFilePath()
	log.Println(saveFilePath)
	SaveJson(coils, saveFilePath)
}
