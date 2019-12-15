package gomon


type coilRequest struct {
	CoilId string      `json:"coilId"`
	CurDir string      `json:"curDir"`
	FatcorNames []string `json:"factorNames"`
}