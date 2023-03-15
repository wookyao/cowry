package filekit

import (
	"os"
	"path/filepath"
	"time"
)

// IsExist 判断文件或者文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return true
}

// IsDir 判断路径是否是目录
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}

	return f.IsDir()
}

// IsFile 判断路径是否是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// MkDir 创建文件夹
func MkDir(path string) (string, error) {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return "", err
	}

	return path, nil
}

// MkDirAll 层级创建所有目录
func MkDirAll(path string, elem ...string) (string, error) {
	paths := []string{path}

	if len(elem) > 0 {
		for _, v := range elem {
			paths = append(paths, v)
		}
	} else {
		paths = append(paths, time.Now().Format("2006/01/02"))
	}

	folderPath := filepath.Join(paths...)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return "", err
	}

	return folderPath, nil
}
