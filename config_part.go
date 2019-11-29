package gomon

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
)

type PartConfig struct {
	Cfg *Config
	Table   string
}

func NewPartConfig(cfg *Config) *PartConfig {
	conf := new(PartConfig)
	conf.Cfg = cfg

	tablePath := fmt.Sprintf(GetComponentsDir() + "/part_table" + conf.Cfg.FmTag +".json")

	log.Println(tablePath)

	byteTable, err := ioutil.ReadFile(tablePath)
	if err != nil {
		panic("Json part table read fail")
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
