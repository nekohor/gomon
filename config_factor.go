package gomon

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"strconv"
)

type FactorConfig struct {
	Cfg *Config
	TomlFactors *TomlFactors

	//TomlSpecifcFactors *TomlSpecificFactors
}

type TomlFactors struct {
	FactorSeries []struct {
		SeriesName  string   `toml:"seriesName"`
		FactorNames []string `toml:"factorNames"`
	} `toml:"factors"`
}

//type TomlSeries struct {
//	SeriesName string `toml:seriesName`
//	FactorNames []string `toml:factorNames`
//}

//type TomlSpecificFactors struct {
//	FactorNames []string `toml:factorNames`
//}

func NewFactorConfig(cfg *Config) *FactorConfig {
	conf := new(FactorConfig)
	conf.Cfg = cfg

	tomlPath := fmt.Sprintf(conf.Cfg.GetComponentsDir() + "/factors.toml")
	if _, err := toml.DecodeFile(tomlPath, &conf.TomlFactors); err != nil {
		CheckError(err)
	}

	log.Println(conf.TomlFactors)
	return conf
}

func (f *FactorConfig) GetFactorNames() []string {
	return f.GetAllFactorNames()
}

func (f *FactorConfig) GetSeriesNames() []string {
	seriesNames := []string{}

	for _, tomlSeries := range f.TomlFactors.FactorSeries {
		seriesNames = append(seriesNames, tomlSeries.SeriesName)
	}
	return seriesNames
}

func (f *FactorConfig) GetAllFactorNames() []string {
	allFactorNames := []string{}

	for _, tomlSeries := range f.TomlFactors.FactorSeries  {
		var factorNamesInSeries []string

		if tomlSeries.SeriesName == "fm_stand" {
			factorNamesInSeries = f.GetFmStandFactorNames(tomlSeries.FactorNames)
		} else {
			factorNamesInSeries = tomlSeries.FactorNames
		}

		for _, factorName := range factorNamesInSeries {
			allFactorNames = append(allFactorNames, factorName)
		}
	}
	log.Println(allFactorNames)
	return allFactorNames
}

func (f *FactorConfig) GetFmStandFactorNames(factorNames []string) []string {
	allFactorNames := []string{}

	for _, std := range GetFmStands() {
		for _, factorName := range factorNames {
			allFactorNames = append(allFactorNames, factorName + strconv.Itoa(std))
		}
	}

	return allFactorNames
}