package gomon

type Part struct {
	size int
	data []dataType
}

func NewPart(ctx *Context, partName string) *Part {
	pt := new(Part)

	line := ctx.GetMillLine()

	dcaFileName := ctx.PartConf.GetDcaFileName(line, partName)
	dcaPath := pt.ConcatPath(ctx.Current.CurDir, ctx.Current.CurCoilId, dcaFileName)

	signalName := ctx.PartConf.GetSignalName(line, partName)

	pt.size, pt.data = ctx.Reader.ReadData(ctx, dcaPath, signalName)

	return pt
}

func (p *Part) ConcatPath(curDir, coilId, dcaFileName string) string {
	return curDir + "/" + coilId + "/" + dcaFileName + ".dca"
}
