package gomon


type Request struct {
	CoilId  string             `json:"coilId"`
	CurDir  string             `json:"curDir"`
	Factors []string           `json:"factors"`
}
