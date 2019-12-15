package gomon


type StatsRequest struct {
	Path Path `json:"path"`
	Stats Stats `json:"stats"`
	LengthDivision LengthDivision `json:"lengthDivision"`
}
type Path struct {
	CurDir string `json:"curDir"`
}
type Stats struct {
	FunctionName string `json:"functionName"`
	Aim int `json:"aim"`
	Tolerance int `json:"tolerance"`
	Upper int `json:"upper"`
	Lower int `json:"lower"`
	Unit string `json:"unit"`
}
type LengthDivision struct {
	LengthName string `json:"lengthName"`
	HeadLen int `json:"headLen"`
	TailLen int `json:"tailLen"`
	HeadCut int `json:"headCut"`
	TailCut int `json:"tailCut"`
}