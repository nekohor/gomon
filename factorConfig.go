package gomon

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"strconv"
)

type FactorConfig struct {
	Setting *Setting
	TomlAll *TomlFactors
	TomlSpecifc TomlSpecificFactors
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

type TomlSpecificFactors struct {
	FactorNames []string `toml:factorNames`
}

func NewFactorConfig(setting *Setting) *FactorConfig {
	conf := new(FactorConfig)
	conf.Setting = setting

	tomlAllPath := fmt.Sprintf(GetComponentsDir() + "/Factors/factors_all.toml")
	tomlSpecificPath := fmt.Sprintf(GetComponentsDir() + "/Factors/factors_specific.toml")

	if _, err := toml.DecodeFile(tomlAllPath, &conf.TomlAll); err != nil {
		CheckError(err)
	}
	if _, err := toml.DecodeFile(tomlSpecificPath, &conf.TomlSpecifc); err != nil {
		CheckError(err)
	}
	log.Println(conf.TomlAll)

	return conf
}

func (f *FactorConfig) GetFactorNames() []string {
	if f.Setting.IsBatchExportMode() && f.Setting.BatchExportMode.IsFactorsSpecific {
		return f.TomlSpecifc.FactorNames
	} else {
		return f.GetAllFactorNames()
	}
}

func (f *FactorConfig) GetSeriesNames() []string {
	seriesNames := []string{}

	for _, tomlSeries := range f.TomlAll.FactorSeries {
		seriesNames = append(seriesNames, tomlSeries.SeriesName)
	}
	return seriesNames
}

func (f *FactorConfig) GetAllFactorNames() []string {
	allFactorNames := []string{}

	for _, tomlSeries := range f.TomlAll.FactorSeries  {
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