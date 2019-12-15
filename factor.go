package gomon

import (
	"log"
)

type Factor struct {
	FactorName string     `json:"factorName"`
	Data       []dataType `json:"data"`
}

func NewFactor(ctx *Context) *Factor {
	factor := new(Factor)
	factor.FactorName = ctx.Current.CurFactorName
	factor.Data = factor.GetData(ctx)
	factor.PrintLog()
	return factor
}

func (f *Factor) PrintLog() {
	log.Println("")

	log.Println(f.FactorName)
	log.Println(len(f.Data))

	printLen := len(f.Data) / 10
	log.Println(f.Data[:printLen])

	log.Println("")
}

func (f *Factor) GetData(ctx *Context) []dataType {
	curData := make([]dataType, 3)
	factorName := ctx.Current.CurFactorName
	switch factorName {
	case "leveling1", "leveling2", "leveling3", "leveling4", "leveling5", "leveling6", "leveling7":
		stdIdx := len(factorName)
		std := string(factorName[stdIdx-1])
		curData = f.GetFactorData2(ctx, "os_gap"+std, "ds_gap"+std)
	case "asym_flt":
		curData = f.GetFactorData2(ctx, "flt_ro1", "flt_ro5")
	case "sym_flt":
		curData = f.GetFactorData3Reverse(ctx, "flt_ro1", "flt_ro3", "flt_ro5")
	case "looper_angle7":
		curData = f.GetFactorData0()
	case "c25_minus_c40":
		curData = f.GetFactorData2(ctx, "crown25", "crown40")
	case "c40_minus_c100":
		curData = f.GetFactorData2(ctx, "crown40", "crown100")
	default:
		// FactorName as partName
		curData = f.GetFactorData1(ctx, factorName)
	}
	return curData
}

func (f *Factor) GetFactorData0() []dataType {
	factorData := make([]dataType, 1)
	return factorData
}

func (f *Factor) GetFactorData1(ctx *Context, single string) []dataType {
	part := NewPart(ctx, single)
	factorData := make([]dataType, part.size)
	for i := 0; i < part.size; i++ {
		factorData[i] = part.data[i]
	}
	return factorData
}

func (f *Factor) GetFactorData2(ctx *Context, os string, ds string) []dataType {
	partOS := NewPart(ctx, os)
	partDS := NewPart(ctx, ds)

	factorData := make([]dataType, partOS.size)
	for i := 0; i < partOS.size; i++ {
		factorData[i] = partOS.data[i] - partDS.data[i]
	}
	return factorData
}

func (f *Factor) GetFactorData3(ctx *Context, os, ct, ds string) []dataType {
	partOS := NewPart(ctx, os)
	partCT := NewPart(ctx, ct)
	partDS := NewPart(ctx, ds)

	factorData := make([]dataType, partCT.size)
	for i := 0; i < partOS.size; i++ {
		factorData[i] = partCT.data[i] - (partOS.data[i]+partDS.data[i])/2
	}
	return factorData

}

func (f *Factor) GetFactorData3Reverse(ctx *Context, os, ct, ds string) []dataType {
	partOS := NewPart(ctx, os)
	partCT := NewPart(ctx, ct)
	partDS := NewPart(ctx, ds)

	factorData := make([]dataType, partCT.size)
	for i := 0; i < partOS.size; i++ {
		factorData[i] = (partOS.data[i]+partDS.data[i])/2 - partCT.data[i]
	}
	return factorData
}
