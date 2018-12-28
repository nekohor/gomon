package gomon


import (
    "os"
    "syscall"
    "unsafe"
    "log"
    "time"
    // "sync"
)

type dataType float32

type DLLCaller struct {
    dllPath string
}

func (d *DLLCaller) pathExists(path string) bool {
    _, err := os.Stat(path)    //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func StrPtr(s string) uintptr {
    return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func INT8FromString(s string) ([]byte, error) {
    for i := 0; i < len(s); i++ {
        if s[i] == 0 {
            return nil, nil
        }
    }
        log.Println(s)
        return []byte(s), nil
    }
    
func StringToINT8(s string) []byte {
    a, err := INT8FromString(s)
    if err != nil {
        panic("syscall: string with NUL passed to StringToINT8")
    }
        // log.Println(a)
        return a
    }

func StringToINT8Ptr(s string) *byte { return &StringToINT8(s)[0] }

func (d *DLLCaller) ReadData(dcaPath, signalName string) (int, []dataType) {
    mydll := syscall.NewLazyDLL(d.dllPath)
    dllReader := mydll.NewProc("ReadData")

    size := 1500
    dataArray := make([]dataType, size)
    if d.pathExists(dcaPath) == true {

        // var wg sync.WaitGroup

        // wg.Add(2)
        // go func() {
        callArgDcaPath := uintptr(unsafe.Pointer(StringToINT8Ptr(dcaPath)))
        callArgSignalName := uintptr(unsafe.Pointer(StringToINT8Ptr(signalName)))

        // wg.Done()

        size_uintptr, _, _ := dllReader.Call(
        callArgDcaPath, callArgSignalName,
        uintptr(unsafe.Pointer(&dataArray[0])))

        // wg.Done()
        size = int(size_uintptr)
        // }()
        time.Sleep(100)
        // wg.Wait()
       
    } else {
        log.Println("dcaPath does not exist: ", dcaPath)
    }
    return size, dataArray
}