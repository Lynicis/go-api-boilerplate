package path

import (
	"path/filepath"
	"runtime"
)

func GetProjectBasePath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	getDirectory := filepath.Join(filepath.Dir(currentFile), "../..")

	return getDirectory
}
