package util

import (
	"os/exec"
	"os"
	"path"
	"path/filepath"
)

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	filePath, _ := filepath.Abs(file)
	dir, _ :=  path.Split(filePath)
	return dir
}
