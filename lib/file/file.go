package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

func CreateDir(path string, perm fs.FileMode) error {
	// filepath.Split 分割路径合文件
	paths, _ := filepath.Split(path)

	if !Exists(paths) {
		err := os.MkdirAll(paths, perm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	//os.Stat获取文件信息
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
