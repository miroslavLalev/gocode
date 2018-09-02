package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const defaultSocketType = "tcp"

var windowsStopSignals = []os.Signal{}

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	proc_get_module_file_name = kernel32.NewProc("GetModuleFileNameW")
)

// Full path of the current executable
func get_executable_filename() string {
	b := make([]uint16, syscall.MAX_PATH)
	ret, _, err := syscall.Syscall(proc_get_module_file_name.Addr(), 3,
		0, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)))
	if int(ret) == 0 {
		panic(fmt.Sprintf("GetModuleFileNameW : err %d", int(err)))
	}
	return syscall.UTF16ToString(b)
}

// Additional OS-specific signals for process termination
func getProcessStopSignals() []os.Signal {
	return windowsStopSignals
}
