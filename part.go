package gomon

type Part struct {
    size int
    data []dataType
}

func NewPart(cfg *Config, partName string) *Part {
    pt := new(Part)

    curDir := cfg.Setting.CurDir
    coilId := cfg.CurCoilId
    line := cfg.Setting.GetMillLine(coilId)

    dcaFileName := cfg.PartTable.GetDcaFileName(line, partName)
    dcaPath := pt.ConcatPath(curDir, coilId, dcaFileName)

    signalName := cfg.PartTable.GetSignalName(line, partName)

    pt.BuildPartData(cfg, dcaPath, signalName)
    return pt
}

func (p *Part) ConcatPath(curDir, coilId, dcaFileName string) string {
    return curDir + "/" + coilId + "/" + dcaFileName + ".dca"
}

func (this *Part) BuildPartData(cfg *Config, dcaPath, signalName string) {
    // this.data = make([]dataType, cfg.Setting.MaxArray)
    this.size, this.data = cfg.DLLCaller.ReadData(cfg, dcaPath, signalName)
    if this.size == -1 {
        this.size = 1
    }
    this.data = this.data[:this.size]
}