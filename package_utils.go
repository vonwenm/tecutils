package tecutils

import (
	"errors"
	"fmt"
	"os"
	"path"
)

//Devuelve la ruta local de un package de go
func GetPackageFullPath(name string) (result string, err error) {
	goPath := os.Getenv("GOPATH")
	result = path.Join(goPath, "src", name)
	if ok := DirectoryExists(result); !ok {
		err = errors.New(fmt.Sprintf("Package %s does not exist", name))
		result = ""
	}

	return
}
