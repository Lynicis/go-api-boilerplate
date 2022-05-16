package project_path

import (
	"path"
	"runtime"
)

func GetRootDirectory() string {
	_, currentFile, _, _ := runtime.Caller(0)
	directory := path.Join(
		path.Dir(currentFile),
		"../..",
	)

	return directory
}
