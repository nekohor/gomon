package gomon

type Coil struct {
	CoilId  string             `json:"coilId"`
	Factors map[string]*Factor `json:"factors"`
}

func NewCoil(ctx *Context, coilId string) *Coil {
	c := new(Coil)
	c.CoilId = coilId

	c.Factors = make(map[string]*Factor)
	factorNames := ctx.FactorConf.GetFactorNames()

	for _, factorName := range factorNames {
		c.Factors[factorName] = NewFactor(ctx, coilId, factorName)
	}
	return c
}
