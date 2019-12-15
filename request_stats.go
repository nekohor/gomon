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
	FunctionName string     `json:"functionName"`
	Aim          float32    `json:"aim"`
	Tolerance    float32    `json:"tolerance"`
	Upper        float32    `json:"upper"`
	Lower        float32    `json:"lower"`
	Unit         string     `json:"unit"`
}
type LengthDivision struct {
	LengthName string `json:"lengthName"`

	HeadLen int `json:"headLen"`
	TailLen int `json:"tailLen"`

	HeadPerc float32 `json:"headPerc"`
	TailPerc float32 `json:"tailPerc"`

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

func (s *StatsRequest) SetPartPerc(headPerc float32, tailPerc float32) {
	s.LengthDivision.HeadPerc = headPerc
	s.LengthDivision.TailPerc = tailPerc
}

func (s *StatsRequest) SetLengthCut(headCut int, tailCut int) {
	s.LengthDivision.HeadCut = headCut
	s.LengthDivision.TailCut = tailCut
}