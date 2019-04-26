package gomon

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type SocketConfig struct {
	Setting *Setting
	Request string
}

func NewSocketConfig(req string) *SocketConfig {
	conf := new(SocketConfig)
	conf.Request = req
	return conf
}

func (s *SocketConfig) GetCoilId() string {
	query := "coilId"
	return gjson.Get(s.Request, query).String()
}

func (s *SocketConfig) GetDcaFileDir() string {
	query := "curDir"
	return gjson.Get(s.Request, query).String()
}

func (s *SocketConfig) GetFactors() []string {
	query := fmt.Sprintf("factors")
	gjsonArray := gjson.Get(s.Request, query).Array()
	stringArray := GJsonArrayToStringArray(gjsonArray)
	return stringArray
}
