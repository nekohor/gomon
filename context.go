package gomon

import (
	// "os"
	// "path/filepath"
	// "log"
	cache "github.com/patrickmn/go-cache"
	"log"
	"time"
)

type Context struct {
	Cfg        *Config
	PartConf   *PartConfig
	FactorConf *FactorConfig

	Reader    *Reader
	CachePool *cache.Cache

	Current *Current
}

type Current struct {
	CurCoilId      string
	CurDir         string
	CurFactorNames []string
	CurFactorName  string
}

func NewContext() *Context {
	this := new(Context)
	this.Reader = NewReader()
	this.Current = &Current{}
	this.CachePool = cache.New(3*time.Minute, 6*time.Minute)
	return this
}

func (c *Context) SetConfig(config *Config) {
	c.Cfg = config
	c.PartConf = NewPartConfig(c.Cfg)
	c.FactorConf = NewFactorConfig(c.Cfg)
}

func (ctx *Context) GetMillLine() string {
	coilId := ctx.Current.CurCoilId
	if string(coilId[0]) == "M" {
		return "1580"
	} else if string(coilId[0]) == "H" {
		return "2250"
	} else if string(coilId[0]) == "G" {
		return "2250"
	} else if string(coilId[0]) == "C" {
		return "1580"
	} else {
		log.Println("In JudgeLine Else Logic")
		log.Println(coilId)
		panic("This coil from wrong line.")
	}
}
