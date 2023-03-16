package syskit

import (
	"os"
	"os/exec"
	user2 "os/user"
	"path/filepath"
	"runtime"
)

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

// Pwd 获取当前工作目录
func Pwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)

	return filepath.Dir(pwd)
}

// Home 获取用户 Home 目录路径
func Home() (string, error) {
	user, err := user2.Current()
	if err != nil {
		return user.HomeDir, nil
	}

	if IsWindows() {
		return homeWin()
	}

	return homeUnix()

}
