package gomon

type StatsRequest struct {
	CoilInfo       CoilInfo       `json:"coilInfo"`
	StatsOption    StatsOption          `json:"stats"`
	LengthDivision LengthDivision `json:"lengthDivision"`
}
type CoilInfo struct {
	CoilId     string `json:"coilId"`
	CurDir  string `json:"curDir"`
	FactorName string `json:"factorName"`
}
type StatsOption struct {
	FunctionName string   `json:"functionName"`
	Aim          DataType `json:"aim"`
	Tolerance    DataType `json:"tolerance"`
	Upper        DataType `json:"upper"`
	Lower        DataType `json:"lower"`
	Unit         string   `json:"unit"`
}
type LengthDivision struct {
	LengthName string `json:"lengthName"`

	HeadLen int `json:"headLen"`
	TailLen int `json:"tailLen"`

	HeadPerc DataType `json:"headPerc"`
	TailPerc DataType `json:"tailPerc"`

	HeadCut int `json:"headCut"`
	TailCut int `json:"tailCut"`
}

func (s *StatsRequest) SetCoilInfo(coilId string, curDir string, factorName string) {
	s.CoilInfo.CoilId = coilId
	s.CoilInfo.CurDir = curDir
	s.CoilInfo.FactorName = factorName
}

func (s *StatsRequest) SetLengthName(name string) {
	s.LengthDivision.LengthName = name
}

func (s *StatsRequest) SetPartLen(headLen int, tailLen int) {
	s.LengthDivision.HeadLen = headLen
	s.LengthDivision.TailLen = tailLen
}

func (s *StatsRequest) SetPartPerc(headPerc DataType, tailPerc DataType) {
	s.LengthDivision.HeadPerc = headPerc
	s.LengthDivision.TailPerc = tailPerc
}

func (s *StatsRequest) SetLengthCut(headCut int, tailCut int) {
	s.LengthDivision.HeadCut = headCut
	s.LengthDivision.TailCut = tailCut
}