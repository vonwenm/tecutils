package tecutils

import (
	"io/ioutil"
	"math"
	"os"
	"path"
	"time"
)

const (
	DEFAULT_PERMISSION = 0777
)

//Returns true if the directory exists
func DirectoryExists(fullpath string) (ok bool) {
	fileInfo, err := os.Lstat(fullpath)
	if err != nil {
		ok = false
		return
	}
	ok = len(fileInfo.Name()) != 0
	return
}

//Verifies if the directory already exists, if it does not then it is created
func CreateDirectoryIfNotExist(fullpath string) (err error) {
	ok := DirectoryExists(fullpath)
	if ok {
		return
	}
	err = os.MkdirAll(fullpath, DEFAULT_PERMISSION)
	return
}

type FileLambda func(fullpath string, f *os.FileInfo)

//Reads a directory content recursively and process the lambda code inside
func ProcessDirectoryContents(fullpath string, recursive bool, fn FileLambda) (err error) {
	files, err := ioutil.ReadDir(fullpath)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() && recursive {
			ProcessDirectoryContents(path.Join(fullpath, file.Name()), recursive, fn)
		}
		if fn != nil {
			fn(fullpath, &file)
		}
	}

	return
}

func FileDaysOld(f *os.FileInfo) int {
	file := *f
	interval := time.Now().Sub(file.ModTime()).Hours() / 24
	return int(math.Floor(interval))
}
