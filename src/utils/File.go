package utils

import "os"

type fileType struct {}

var File fileType

/**
文件是否存在
 */
func (this fileType) Exist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
