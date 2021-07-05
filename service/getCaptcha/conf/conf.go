package conf

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetCurrentAbS() string{
	var BaseDir string
	_, filename,_,ok := runtime.Caller(0)
	if ok {
		AbsPath := path.Dir(filename)
		BaseDir = filepath.Dir(AbsPath)
	}
	return BaseDir
}







