package gomon

import "bytes"

type Stats struct {
	ctx *Context

	rawData []dataType
	data []dataType
	dataSize int
	req *StatsRequest
}

func NewStats(ctx *Context, data []dataType, req *StatsRequest) *Stats {
	stats :=  &Stats{}
	stats.ctx = ctx

	stats.rawData = data
	stats.data = data
	stats.dataSize = len(data)

	stats.req = req

	return stats
}



func (s *Stats) GetHeadTailCut() (int, int) {

	headCut := s.req.LengthDivision.HeadCut
	tailCut := s.req.LengthDivision.TailCut

	return headCut, tailCut

}

func (s *Stats) GetHeadTailLen() (int, int) {

	var headLen int
	var tailLen int

	if s.req.LengthDivision.HeadPerc < 0 || s.req.LengthDivision.TailPerc < 0 {
		headLen = s.req.LengthDivision.HeadLen
		tailLen = s.req.LengthDivision.TailLen
	} else {
		headLen = int(s.req.LengthDivision.HeadPerc * float32(s.dataSize))
		tailLen = int(s.req.LengthDivision.TailPerc * float32(s.dataSize))
	}
	return headLen, tailLen
}

func (s *Stats) GetIndex() (int, int) {

	var startIndex int
	var endIndex int
	lastIndex := s.dataSize - 1


	headCut, tailCut := s.GetHeadTailCut()
	headLen, tailLen := s.GetHeadTailLen()

	switch s.req.LengthDivision.LengthName {
	case "total":
		startIndex = 0
		endIndex = lastIndex
	case "main":
		startIndex = headCut
		endIndex = lastIndex - tailCut
	case "head":
		startIndex = headCut
		endIndex = headCut + headLen
	case "middle":
		startIndex = headCut + headLen
		endIndex = lastIndex - ( tailCut + tailLen )
	case "tail":
		startIndex = lastIndex - ( tailCut + tailLen )
		endIndex = lastIndex - tailCut
	case "body":
		startIndex = headCut + headLen
		endIndex = lastIndex - tailCut
	case "first":
		startIndex = headCut
		endIndex = headCut + 5
	default:
		//	headLen tailLen as startLen and endLen with 0
		x1Len := headLen
		x2Len := tailLen
		startIndex = headCut + x1Len
		endIndex = headCut + x2Len
	}

	if startIndex < 0 || startIndex >= s.dataSize {
		startIndex, endIndex = 0, 0
	}
	if endIndex < 0 || endIndex >= s.dataSize {
		startIndex, endIndex = 0, 0
	}

	return startIndex, endIndex
}

func (s *Stats) SelectLength() []dataType {
	startIndex, endIndex := s.GetIndex()

	return s.data[startIndex : endIndex]
}

func (s *Stats) Calculate() float32 {

	s.data = s.SelectLength()
	switch s.req.StatsOption.FunctionName {
	case "aimRate":
		s.CalcAimRate()
	}

}


func (s *Stats) GetTolerance() (float32, float32){
	if s.req.StatsOption.Tolerance == 0 {

	}
}

func (s *Stats) CalcAimRate() {


}