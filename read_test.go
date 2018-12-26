package gomon

import (
    "testing"
    "log"
)

func TestReadData(t *testing.T) {
    // fileName := "d:/test/H18090482A/FDT_POND.dca"
    fileName := "d:/test/H181657280/FDT_POND.dca"
    // fileName := "d:/test/H181657280/CLG_POND.dca"
    signalName := "TN\\L_FA_FDT1TEMP"
    // signalName := "TN\\L_FM_FDTTRANSLEN"
    // signalName := "TNN\\L_FA_FDT1TEMP"
    // signalName := "FDTN\\Calc_Length"
    // signalName := "TN\\L_AG2_F7XTHKCDEVCLG"
    log.Println("    test fileName ", fileName)
    log.Println("    test signalName ", signalName)
    var size int
    var data = make([]dataType, 1500)
    DLLCaller := &DLLCaller{"C:/Users/Crystal/go/src/github.com/nekohor/gomon/app" + "/ReadDCADLL.dll"}
    size, data = DLLCaller.ReadData(fileName, signalName)
    log.Println(data[0:25])
    log.Println(size)
}