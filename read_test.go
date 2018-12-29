package gomon

import (
    "testing"
    "log"
)

func TestReadData(t *testing.T) {
    // fileName := "d:/test/H18090482A/FDT_POND.dca"
    fileName := "d:/test/H181657280/FDT_POND.dca"
    // fileName := "E:/1580hrm/201811/20181101.zip/20181101/M18132136H/R2DW_POND.dca"
    // fileName := "d:/test/H181657280/CLG_POND.dca"
    
    // signalName := "TN\\L_FA_FDT1TEMP"
    // signalName := "R2DW\\Calc_Length"
    // signalName := "TN\\L_FM_FDTTRANSLEN"
    signalName := "TN\\L_FA_FDT1TEMP"
    // signalName := "FDT\\Calc_Length"
    // signalName := "TN\\L_AG2_F7XTHKCDEVCLG"

    // log.Println("    test fileName ", fileName)
    // log.Println("    test signalName ", signalName)
    var size int
    var data = make([]dataType, 2000)
    DLLCaller := &DLLCaller{"C:/Users/Crystal/go/src/github.com/nekohor/gomon/app" + "/ReadDCADLL.dll"}
    size, data = DLLCaller.ReadData(fileName, signalName)
    log.Println(data[0:25])
    log.Println(size)
}