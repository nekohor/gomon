# Golang调用dll函数的问题

各位大佬好！

## 问题描述

各位大佬好！我用golang调用dll函数遇到问题。dll函数的参数类型是`const char*`，单次调用没有问题，多次调用时在dll接收到的字符串会偶尔出现字符串增多的现象。

### dll函数情况

```c++
extern "C" ReadDCADLL_API 
long ReadData(const char* DcaFile, const char* SignalName, float* DataArray);
```

其中DcaFile是出问题的字符串，`DcaFile`是文件的路径，`ReadData`函数通过`DcaFile`读取文件并将相应的信号读取出来保存到DataArray数组里。

### 调用过程

golang调用dll函数如下

```go
mydll, err := syscall.LoadLibrary("xxx.dll")
dllReader, err := syscall.GetProcAddress(mydll, "ReadData")

// 参数类型均经过转换
callArgDcaPath = uintptr(unsafe.Pointer(StringToINT8Ptr(dcaPath)))
callArgSignalName = uintptr(unsafe.Pointer(StringToINT8Ptr(signalName)))
dumpArray = uintptr(unsafe.Pointer(&dataArray[0]))

// golang调用dll函数
size_uintptr, _, _ = syscall.Syscall(dllReader,3,
            callArgDcaPath, callArgSignalName, dumpArray)
```

### 单次调用无问题

以上过程单次调用无问题，可以读取到数据。

### 多次调用存在问题

循环多次调用的话，偶尔会出现golang传递的参数和dll接收到的参数不一致的问题。golang对应的`callArgDcaPath`字符串是正确的，但是dll函数接收的`DcaFile`尾部随机多出字符串。

```go
// 正确的字符串
D:/Work/sample/data/pOND/H19013640A/FET_POND.dca

//随机增多后的字符串有时为：
D:/Work/sample/data/pOND/H19013640A/FET_POND.dcaD:/Work/sample/data/pOND/H19013638L/F2_POND.dca'

//随机增多后的字符串有时为：
D:/Work/sample/data/pOND/H19013640A/FET_POND.dcax\x01
```

请问各位大佬这种问题如何解决？



各位大佬好！我用golang调用dll函数遇到问题。dll函数的参数类型是const char*，单次调用没有问题，循环多次调用时在dll接收到的字符串会偶尔出现字符串增多的现象。尾部多出来的字符有时候是\x01,有时候是上一个循环的字符串，这一般是怎么回事？