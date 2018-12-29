package gomon

import (
    "io/ioutil"
    "fmt"
    "github.com/tidwall/gjson"
)

type PartTable struct {
    table string
}

type FactorTable struct {
    table string
    factorIds []string
}

func ReadTable(exeDir,tableName string) string {
    tableFilePath := fmt.Sprintf("/tables/%s_table.json",tableName)
    table, err := ioutil.ReadFile(exeDir + tableFilePath)
    if err != nil {
        panic("part table json setting read fail")
    }
    return string(table)
}

func NewPartTable(exeDir string) *PartTable {
    pt := new(PartTable)
    pt.table = ReadTable(exeDir, "part")
    return pt
}

func (pt *PartTable) GetDcaFileName(line string, partName string) string {
    queryDcaFileName := fmt.Sprintf("part_table.#[line==%s].table.#[part==%s].dcafile", line, partName)
    return gjson.Get(pt.table, queryDcaFileName).String()
}

func (pt *PartTable) GetSignalName(line string, partName string) string {
    querySignalName := fmt.Sprintf("part_table.#[line==%s].table.#[part==%s].signal", line, partName)
    return gjson.Get(pt.table, querySignalName).String()
}

func NewFactorTable(exeDir string) *FactorTable {
    ft := new(FactorTable)
    ft.table = ReadTable(exeDir, "factor")
    ft.factorIds = ft.GetFactorIds()
    return ft
}

func (ft *FactorTable) GetSeriesArray() []string {
    query := fmt.Sprintf("factor_table.#.seriesName")
    gjsonArray := gjson.Get(ft.table, query).Array()
    stringArray := ResultArrayToStringArray(gjsonArray)
    return stringArray
}

func (ft *FactorTable) GetFactorIds() []string {
    allFactorIds := []string{}
    seriesList := ft.GetSeriesArray()

    for _, seriesName := range seriesList {
        query := fmt.Sprintf("factor_table.#[seriesName==%s].factorList",seriesName)
        gjsonArray := gjson.Get(ft.table, query).Array()
        stringArray := ResultArrayToStringArray(gjsonArray)

        for _, taskName := range stringArray {
            allFactorIds = append(allFactorIds, taskName)
        }
    }
    return allFactorIds
}

func (ft *FactorTable) GetFactorNameZhCn(factorName string) string {
    query := fmt.Sprintf("factor_table.#[seriesName==%s].factorList",factorName)
    return  gjson.Get(ft.table, query).String()
}



func ResultArrayToStringArray(resArr []gjson.Result) []string {
    strArr := make([]string, len(resArr))
    for i, v := range resArr {
        strArr[i] = v.String()
    }
    return strArr
}