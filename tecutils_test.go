package tecutils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"testing"
)

func TestUUIDNoDuplicates(t *testing.T) {
	iterations := 50000
	log.Printf("Processing %v iterations", iterations)
	values := make(map[string]bool)
	dups := 0

	for i := 0; i < iterations; i++ {
		v := UUID()
		if _, ok := values[v]; ok {
			dups++
		}
		values[v] = true
	}
	assert.Equal(t, dups, 0)
}

func TestEncrypt(t *testing.T) {
	pwd := UUID()
	encripted := Encrypt(pwd)
	log.Printf("Original string is %v long. Encripted string is %v", len(pwd), len(encripted))
	assert.NotEqual(t, pwd, encripted)
}

func TestUrlParser(t *testing.T) {
	url := "https://localhost:8080/mau/testing?something=ok"
	parsed, err := ParseBaseUrl(url)
	log.Println(parsed)
	assert.NoError(t, err)
	assert.Equal(t, parsed, "https://localhost:8080")
}

func TestPackageFullPath(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	packageName := "github.com/mauleyzaola/tecutils"
	fullPath := path.Join(goPath, "src", packageName)
	result, err := GetPackageFullPath(packageName)
	assert.Equal(t, result, fullPath)
	assert.NoError(t, err)
	log.Println("this package's path is:", result)
}

func TestExecutableDirectory(t *testing.T) {
	curr, err := GetExecutableDirectory()
	assert.NoError(t, err)
	log.Println("Current directory is:", curr)
}
