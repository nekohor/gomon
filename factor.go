package gomon

import (
    "log"
)
type Factor struct {
    FactorName string `json:"factorName"`
    FactorNameZhCn string `json:"factorNameZhCn"`
    Data []dataType `json:"data"`
}

func NewFactor(cfg *Config) *Factor {
    factor := new(Factor)
    factor.FactorName =  cfg.CurFactorName
    // factor.FactorNameZhCn = cfg.FactorTable.GetFactorNameZhCn(cfg.CurFactorName)
    factor.BuildData(cfg)
    return factor
}

func (this *Factor) BuildData(cfg *Config) {
    switch cfg.CurFactorName {
        case "leveling1", "leveling2", "leveling3", "leveling4", "leveling5", "leveling6", "leveling7":
            std_idx := len(cfg.CurFactorName)
            std := string(cfg.CurFactorName[std_idx - 1])
            this.BuildFactorData2(cfg,"os_gap" + std,"ds_gap" + std)
        case "asym_flt":
            this.BuildFactorData2(cfg,"flt_ro1","flt_ro5")
        case "sym_flt":
            this.BuildFactorData3(cfg,"flt_ro1","flt_ro3","flt_ro5")
        case "looper_angle7":
            this.BuildFactorData0()
        default:
            // FactorName as partName
            this.BuildFactorData1(cfg,cfg.CurFactorName)
    }
}

func (this *Factor) BuildFactorData0() {
    this.Data = make([]dataType, 1)
}

func (this *Factor) BuildFactorData1(cfg *Config, partName string) {
    p := NewPart(cfg, partName)
    log.Println(p)
    this.Data = make([]dataType, p.size)
    for i := 0; i < p.size; i++ {
        this.Data[i] = p.data[i]
    }
}

func (t *Factor) BuildFactorData2(cfg *Config, os string, ds string) {
    p_os := NewPart(cfg, os)
    p_ds := NewPart(cfg, ds)

    t.Data = make([]dataType, p_os.size)
    for i := 0; i < p_os.size; i++ {
        t.Data[i] = p_os.data[i] - p_ds.data[i]
    }
}

func (t *Factor) BuildFactorData3(cfg *Config, os, ct, ds string) {
    p_os := NewPart(cfg, os)
    p_ct := NewPart(cfg, ct)
    p_ds := NewPart(cfg, ds)

    t.Data = make([]dataType, p_ct.size)
    for i := 0; i < p_os.size; i++ {
        t.Data[i] = (p_os.data[i] + p_ds.data[i]) / 2 - p_ct.data[i]
    }
}

