package gomon

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type SocketConfig struct {

	ReqString string

	//Setting *Setting
	CoilIds []string
	Requests map[string]*Request

}

func NewSocketConfig(req string) *SocketConfig {
	conf := new(SocketConfig)
	conf.ReqString = req

	return conf
}

func (s *SocketConfig) GetCoilIds() []string {
	query := "requests.#.coilId"
	gjsonArray := gjson.Get(s.ReqString, query).Array()
	stringArray := GJsonArrayToStringArray(gjsonArray)
	return stringArray
}

func (s *SocketConfig) GetCurDir(coilId string) string {
	query := fmt.Sprintf("requests.#[coilId==\"%s\"].curDir", coilId)
	return gjson.Get(s.ReqString, query).String()
}

func (s *SocketConfig) GetFactors(coilId string) []string {
	query := fmt.Sprintf("requests.#[coilId==\"%s\"].factors", coilId)
	gjsonArray := gjson.Get(s.ReqString, query).Array()
	stringArray := GJsonArrayToStringArray(gjsonArray)
	return stringArray
}
