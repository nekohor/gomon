package gomon




type Part struct {
    size int
    data []dataType
}


func NewPart(cfg *Config, partName string) *Part {
    pt := new(Part)

    curDir := cfg.PathConfig.CurDir
    coilId := cfg.CurCoilId
    line := pt.JudgeLine(coilId)

    dcaFileName := cfg.PartTable.GetDcaFileName(line, partName)
    dcaPath := pt.ConcatPath(curDir, coilId, dcaFileName)

    signalName := cfg.PartTable.GetSignalName(line, partName)

    pt.BuildPartData(cfg, dcaPath, signalName)
    return pt
}

func (this *Part) JudgeLine(coilId string) int {
    if string(coilId[0]) == "M" {
        return 1580
    } else if string(coilId[0]) == "H" {
        return 2250
    } else {
        panic("This coil from wrong line.")
    }
}

func (p *Part) ConcatPath(curDir, coilId, dcaFileName string) string {
    return curDir + "/" + coilId + "/" + dcaFileName + ".dca"
}

func (this *Part) BuildPartData(cfg *Config, dcaPath, signalName string) {
    this.data = make([]dataType, 1500)
    this.size, this.data = cfg.DLLCaller.ReadData(dcaPath, signalName)
    if this.size == -1 {
        this.size = 50
    }
}