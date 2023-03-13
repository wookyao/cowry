package syskit

import "runtime"

// IsWindows 宿主环境 windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux 宿主环境 linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac 宿主环境 mac
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsBits 获取系统的字节
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}
