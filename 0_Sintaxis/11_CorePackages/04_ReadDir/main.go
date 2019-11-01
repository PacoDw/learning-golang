package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ReadDirLongerWay(".")
	ReadDirShorterWay(".")
}

//ReadDirLongerWay To read File longer way...
func ReadDirLongerWay(path string) {
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

//ReadDirShorterWay The shorter way
func ReadDirShorterWay(myPath string) {
	filepath.Walk(myPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
