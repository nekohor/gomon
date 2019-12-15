package gomon

type Coil struct {
	CoilId  string             `json:"coilId"`
	Factors map[string]*Factor `json:"factors"`
}

func NewCoil(ctx *Context) *Coil {

	c := new(Coil)
	c.CoilId = ctx.Current.CurCoilId
	c.Factors = make(map[string]*Factor)
	for _, factorName := range ctx.Current.CurFactorNames {
		ctx.Current.CurFactorName = factorName
		c.Factors[factorName] = NewFactor(ctx)
	}
	return c
}
