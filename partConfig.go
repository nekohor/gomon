package gomon

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

type PartConfig struct {
	Setting *Setting
	Table   string
}

func NewPartConfig(setting *Setting) *PartConfig {
	conf := new(PartConfig)
	conf.Setting = setting

	tablePath := fmt.Sprintf(GetComponentsDir() + "/Tables/partTable.json")
	byteTable, err := ioutil.ReadFile(tablePath)
	if err != nil {
		panic("Json partTable read fail")
	}
	conf.Table = string(byteTable)

	return conf
}

func (p *PartConfig) GetDcaFileName(line string, partName string) string {
	queryDcaFileName := fmt.Sprintf("partTable.#[line==%s].table.#[part==%s].dcafile", line, partName)
	return gjson.Get(p.Table, queryDcaFileName).String()
}

func (p *PartConfig) GetSignalName(line string, partName string) string {
	querySignalName := fmt.Sprintf("partTable.#[line==%s].table.#[part==%s].signal", line, partName)
	return gjson.Get(p.Table, querySignalName).String()
}
