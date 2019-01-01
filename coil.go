package gomon

import (
    "sync"
)

type Coil struct {
    CoilId string `json:"coilId"`
    Factors map[string]*Factor `json:"factors"`
}


func NewCoil(coilId string) *Coil {
    coil := new(Coil)
    coil.CoilId = coilId
    coil.Factors = make(map[string]*Factor)
    return coil 
} 

func (this *Coil) PutData(cfg *Config) {
    cfg.CurCoilId = this.CoilId
    factorArray := cfg.Setting.GetFactorArray(cfg.FactorTable.factorIds)
    for _, factorName := range factorArray {
        cfg.CurFactorName = factorName
        var l *sync.Mutex
        l = new(sync.Mutex)
        l.Lock()
        defer l.Unlock()
        this.Factors[factorName] = NewFactor(cfg)
    }
}

