package main

import "syscall"

var (
	// Syscall to kernel32.dll for isDebuggerPresent
	isDebuggerPresent = syscall.NewLazyDLL("kernel32.dll").NewProc("IsDebuggerPresent")
)

func checkPresents() bool {
	// Check for debugger
	flag, _, _ := isDebuggerPresent.Call()
	if flag != 0 {
		return true
	}
	return false
}
