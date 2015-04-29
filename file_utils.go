package tecutils

import (
	"github.com/kardianos/osext"
	"io/ioutil"
	"os"
	"strings"
)

type matchedFiles struct {
	File  string `json:"fileName"`
	Theme string `json:"theme"`
}

func GetFiles(dirname, suffix string) []matchedFiles {
	files, _ := ioutil.ReadDir(dirname)
	res := make([]matchedFiles, 0)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) && len(f.Name()) != 0 && !f.IsDir() {
			var item matchedFiles
			item.File = f.Name()
			item.Theme = strings.Replace(item.File, suffix, "", -1)
			res = append(res, item)
		}
	}
	return res
}

func DirectoryExists(filePath string) (ok bool) {
	_, err := os.Stat(filePath)
	if err == nil {
		ok = true
	}
	if os.IsNotExist(err) {
		ok = false
	}
	return
}

func GetExecutableDirectory() (string, error) {
	return osext.ExecutableFolder()
}
