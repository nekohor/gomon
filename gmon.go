package gomon

import (
	// "os"
	// "path/filepath"
	"log"
)

type GoMonitor struct {
	Context *Context
}

func NewGoMonitor() *GoMonitor {
	g := new(GoMonitor)
	g.Context = NewContext()
	return g
}

func (g *GoMonitor) GetCoil(coilId string) *Coil {
	coil := NewCoil(g.Context, coilId)
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
	coils = g.GetCoils(g.Context.CoilIds)
	saveFilePath := g.Context.Setting.GetResultFilePath()
	log.Println(saveFilePath)
	SaveJson(coils, saveFilePath)
}
