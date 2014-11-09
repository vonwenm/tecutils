package tecutils

import (
	"io/ioutil"
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
