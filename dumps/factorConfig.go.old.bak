package gomon

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

type FactorConfig struct {
	Setting *Setting
	Table   string
}

func NewFactorConfig(setting *Setting) *FactorConfig {
	conf := new(FactorConfig)
	conf.Setting = setting

	tablePath := fmt.Sprintf(GetComponentsDir() + "/Tables/factorTable.json")
	byteTable, err := ioutil.ReadFile(tablePath)
	if err != nil {
		panic("Json factorTable read fail")
	}
	conf.Table = string(byteTable)

	return conf
}

func (f *FactorConfig) GetFactorNames() []string {
	if f.Setting.SpecificFactorsMode {
		return f.Setting.SpecificFactors
	} else {
		return f.GetAllFactorNames()
	}
}

func (f *FactorConfig) GetSeriesArray() []string {
	query := fmt.Sprintf("factorTable.#.seriesName")
	gjsonArray := gjson.Get(f.Table, query).Array()
	stringArray := GJsonArrayToStringArray(gjsonArray)
	return stringArray
}

func (f *FactorConfig) GetAllFactorNames() []string {
	allFactorIds := []string{}
	seriesList := f.GetSeriesArray()

	for _, seriesName := range seriesList {
		query := fmt.Sprintf("factorTable.#[seriesName==%s].taskList", seriesName)
		resultArray := gjson.Get(f.Table, query).Array()
		stringArray := GJsonArrayToStringArray(resultArray)

		for _, taskName := range stringArray {
			allFactorIds = append(allFactorIds, taskName)
		}
	}
	return allFactorIds
}
