package gomon

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type RequestConfig struct {

	ReqString string

	CoilIds []string
	Requests map[string]*Request

}

func NewRequestConfig(req string) *RequestConfig {
	conf := new(RequestConfig)
	conf.ReqString = req

	return conf
}

func (s *RequestConfig) GetCoilIds() []string {
	query := "requests.#.coilId"
	gjsonArray := gjson.Get(s.ReqString, query).Array()
	stringArray := GJsonArrayToStringArray(gjsonArray)
	return stringArray
}

func (s *RequestConfig) GetCurDir(coilId string) string {
	query := fmt.Sprintf("requests.#[coilId==\"%s\"].curDir", coilId)
	return gjson.Get(s.ReqString, query).String()
}

func (s *RequestConfig) GetFactors(coilId string) []string {
	query := fmt.Sprintf("requests.#[coilId==\"%s\"].factors", coilId)
	gjsonArray := gjson.Get(s.ReqString, query).Array()
	stringArray := GJsonArrayToStringArray(gjsonArray)
	return stringArray
}
