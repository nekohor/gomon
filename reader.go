package gomon

/*
#include <stdio.h>   // 如果要调用C.free 一定要在些包含对应的头文件
#include <stdlib.h>
*/
import "C"
import (
	"log"
	//"os"
	"syscall"
	"unsafe"

	// "time"
	"sync"
)

type dataType float32

type Reader struct {
	readFunc uintptr
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
	return reader
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

func (reader *Reader) ReadData(ctx *Context, dcaPath, signalName string) (int, []dataType) {

	size := ctx.Setting.DataMaxNum
	dataArray := make([]dataType, size)
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

		var l *sync.Mutex
		l = new(sync.Mutex)
		l.Lock()
		sizeUintptr, _, _ := syscall.Syscall(
			reader.readFunc, 3,
			callArgDcaPath,
			callArgSignalName,
			callArgDataArray)
		l.Unlock()

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
	return size, dataArray
}
