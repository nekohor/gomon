package gomon

/*
#include <stdio.h>   // 如果要调用C.free 一定要在些包含对应的头文件
#include <stdlib.h>  // 此段注释与 import "C" 之间不能有空格
*/
import "C"
import (
	"log"
	"math"

	//"os"
	"syscall"
	"unsafe"

	// "time"
	"sync"
)

type DataType float32

type Reader struct {
	readFunc uintptr

	mutex *sync.Mutex
}

func NewReader() *Reader {
	reader := new(Reader)
	handle, err := syscall.LoadLibrary("ReadDCADLL.dll")
	if err != nil {
		log.Println(err)
		panic("err in LoadLibrary")
	}
	//defer syscall.FreeLibrary(handle)

	readFunc, err := syscall.GetProcAddress(handle, "ReadData")
	if err != nil {
		panic("err in GetProcAddress")
	}
	reader.readFunc = readFunc

	reader.mutex = new(sync.Mutex)

	return reader
}


func (reader *Reader) ReadData(ctx *Context, dcaPath, signalName string) (int, []DataType) {

	reader.mutex.Lock()
	defer reader.mutex.Unlock()

	size := ctx.Cfg.DataMaxNum
	dataArray := make([]DataType, size)
	if IsExist(dcaPath) == true {

		//callArgDcaPath := uintptr(unsafe.Pointer(StringToINT8Ptr(dcaPath)))
		//callArgSignalName := uintptr(unsafe.Pointer(StringToINT8Ptr(signalName)))

		gbkDcaPath, err := Utf8ToGbk([]byte(dcaPath))
		if err != nil {
			log.Println(err)
		}

		gbkSignalName, err := Utf8ToGbk([]byte(signalName))
		if err != nil {
			log.Println(err)
		}

		CgoDcaPath := C.CString(string(gbkDcaPath))
		CgoSignalName := C.CString(string(gbkSignalName))
		defer C.free(unsafe.Pointer(CgoDcaPath))
		defer C.free(unsafe.Pointer(CgoSignalName))

		callArgDcaPath := uintptr(unsafe.Pointer(CgoDcaPath))
		callArgSignalName := uintptr(unsafe.Pointer(CgoSignalName))
		callArgDataArray := uintptr(unsafe.Pointer(&dataArray[0]))

		sizeUintptr, _, _ := syscall.Syscall(
			reader.readFunc, 3,
			callArgDcaPath,
			callArgSignalName,
			callArgDataArray)

		size = int(sizeUintptr)

	} else {
		log.Println("dcaPath does not exist: ", dcaPath)
		size = -2
	}

	log.Println("actual return size")
	log.Println(size)

	if -1 == size || -2 == size {
		size = 1
		log.Println("[Warning] wrong DCA path or signal name in DLL function")
	}

	buffArray := make([]DataType, len(dataArray))

	for i := 0; i < len(dataArray); i++ {

		if math.IsNaN(float64(dataArray[i])) {
			buffArray[i] = 0
		} else {
			buffArray[i] = dataArray[i]
		}
	}

	return size, buffArray
}


// ==================== without cgo start ======================
//func StrPtr(s string) uintptr {
//	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
//}
//
//func INT8FromString(s string) ([]byte, error) {
//	for i := 0; i < len(s); i++ {
//		if s[i] == 0 {
//			return nil, nil
//		}
//	}
//	log.Println(s)
//	return []byte(s), nil
//}
//
//func StringToINT8(s string) []byte {
//	a, err := INT8FromString(s)
//	if err != nil {
//		panic("syscall: string with NULL passed to StringToINT8")
//	}
//	// log.Println(a)
//	return a
//}
//
//func StringToINT8Ptr(s string) *byte { return &StringToINT8(s)[0] }
// ========================== without cgo end ===========================