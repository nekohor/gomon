package gomon

import (
	"fmt"
	"math"
)

type Stats struct {
	ctx *Context
	data []DataType
	req *StatsRequest
}

func NewStats(ctx *Context, data []DataType, req *StatsRequest) *Stats {
	stats :=  &Stats{}
	stats.ctx = ctx
	stats.data = data
	stats.req = req

	return stats
}

func (s *Stats) GetDataSize() int {
	return len(s.data)
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
		dataSize := s.GetDataSize()
		headLen = int(s.req.LengthDivision.HeadPerc * DataType(dataSize))
		tailLen = int(s.req.LengthDivision.TailPerc * DataType(dataSize))
	}
	return headLen, tailLen
}

func (s *Stats) GetIndex() (int, int) {

	var startIndex int
	var endIndex int
	lastIndex := s.GetDataSize() - 1


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

	if startIndex < 0 || startIndex > lastIndex {
		startIndex, endIndex = 0, 0
	}
	if endIndex < 0 || endIndex > lastIndex {
		startIndex, endIndex = 0, 0
	}

	return startIndex, endIndex
}

func (s *Stats) SelectLength() []DataType {
	startIndex, endIndex := s.GetIndex()

	return s.data[startIndex : endIndex]
}

func (s *Stats) Calculate() string {

	dataForCalc := s.SelectLength()

	var result string

	switch s.req.StatsOption.FunctionName {
	case "aimRate":
		result = Round(s.CalcAimRate(dataForCalc))
	case "stability":
		result = Round(s.CalcStability(dataForCalc))
	case "mean":
		result = Round(s.CalcMean(dataForCalc))
	case "absMean":
		result = Round(s.CalcAbsMean(dataForCalc))
	case "max":
		result = Round(s.CalcMax(dataForCalc))
	case "absMax":
		result = Round(s.CalcAbsMax(dataForCalc))
	case "min":
		result = Round(s.CalcMin(dataForCalc))
	case "absMin":
		result = Round(s.CalcAbsMin(dataForCalc))
	case "std":
		result = Round(s.CalcStd(dataForCalc))
	case "pickWedge":
		result = s.CalcPickWedge(dataForCalc)
	default:
		result = ""
	}

	return result

}

func Round(val DataType) string {
	return fmt.Sprintf("%.3f", val)
}

func GetSize(data []DataType) int {
	return len(data)
}

func Abs(val DataType) DataType {

	var absVal DataType
	if val < 0 {
		absVal = - val
	} else {
		absVal = val
	}
	return absVal
}

func (s *Stats) GetUpperLower() (DataType, DataType){

	var upper DataType
	var lower DataType

	if s.req.StatsOption.Tolerance == 0 {

		upper = s.req.StatsOption.Upper
		lower = s.req.StatsOption.Lower

	} else {
		aim := s.req.StatsOption.Aim
		tol := s.req.StatsOption.Tolerance

		upper = aim + tol
		lower = aim - tol
	}

	return s.convertUnits(upper), s.convertUnits(lower)
}


func (s *Stats) convertUnits(val DataType) DataType {

	var converted DataType
	if s.req.StatsOption.Unit == "um" {
		converted = val / 1000
	} else {

	}
	return converted
}


func (s *Stats) CalcMean(data []DataType) DataType {

	var sum DataType = 0
	for i := 0;i < len(data) ; i++  {
		sum += data[i]
	}

	return DataType(sum) / DataType(GetSize(data))

}

func (s *Stats) CalcAbsMean(data []DataType) DataType {

	var sum DataType = 0
	for i := 0; i < len(data) ; i++  {
		sum += Abs(data[i])
	}
	return DataType(sum) / DataType(GetSize(data))

}

func (s *Stats) CalcMax(data []DataType) DataType {

	var maxVal DataType = 0
	for i := 0; i < len(data) ; i++  {
		if maxVal < data[i] {
			maxVal = data[i]
		}
	}
	return maxVal

}

func (s *Stats) CalcAbsMax(data []DataType) DataType {

	var maxVal DataType = 0
	for i := 0; i < len(data) ; i++  {
		if maxVal < Abs(data[i]) {
			maxVal = Abs(data[i])
		}
	}
	return maxVal

}

func (s *Stats) CalcMin(data []DataType) DataType {

	var minVal DataType = 0
	for i := 0; i < len(data) ; i++  {
		if minVal > data[i] {
			minVal = data[i]
		}
	}
	return minVal

}

func (s *Stats) CalcAbsMin(data []DataType) DataType {

	var minVal DataType = 0
	for i := 0; i < len(data) ; i++  {
		if minVal > Abs(data[i]) {
			minVal = Abs(data[i])
		}
	}
	return minVal

}


func (s *Stats) CalcAimRate(data []DataType) DataType {

	upper, lower := s.GetUpperLower()

	shippedSize := 0
	for i := 0;i < len(data) ; i++  {
		if data[i] <= upper && data[i] >= lower {
			shippedSize ++
		}
	}
	return DataType(shippedSize) / DataType(len(data)) * 100
}


func (s *Stats) CalcStability(data []DataType) DataType {

	mean := s.CalcMean(data)
	upper := mean + s.convertUnits(s.req.StatsOption.Tolerance)
	lower := mean - s.convertUnits(s.req.StatsOption.Tolerance)

	shippedSize := 0
	for i := 0;i < len(data) ; i++  {
		if data[i] <= upper && data[i] >= lower {
			shippedSize ++
		}
	}
	return DataType(shippedSize) / DataType(len(data)) * 100
}

func (s *Stats) CalcStd(data []DataType) DataType {

	//数学期望
	var sum DataType = 0
	for _, v := range data {
		sum += v
	}
	u := DataType(sum) / DataType(len(data))

	//标准差
	var variance DataType
	for _, v := range data {
		variance += DataType( math.Pow(float64(v - u), 2) )
	}
	sigma := DataType(math.Sqrt(float64(variance) / float64(len(data))))

	return sigma

}


func (s *Stats) CalcPickWedge(data []DataType) string {

	mid := int(len(data) / 2)

	cut := 45
	tol := DataType(0.02)
	firstChild := data[cut: mid - 1]
	secondChild := data[mid: len(data) - 1 - cut]

	var firstPass int
	var secondPass int
	if s.CalcMax(firstChild) - s.CalcMin(firstChild) <= tol {
		firstPass = 1
	} else {
		firstPass = 0
	}

	if s.CalcMax(secondChild) - s.CalcMax(secondChild) <= tol {
		secondPass = 1
	} else {
		secondPass = 0
	}

	return fmt.Sprintf("%d|%d", firstPass, secondPass)

}