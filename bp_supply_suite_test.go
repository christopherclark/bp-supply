package main_test

import (
	"archive/zip"
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBpSupply(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BpSupply Suite")
}

func ZipFileContent(zipFile, path string) (string, error) {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}
	defer r.Close()

	for _, zf := range r.File {
		if zf.Name == path {
			rc, err := zf.Open()
			if err != nil {
				return "", fmt.Errorf("%s: open compressed file: %v", zf.Name, err)
			}
			defer rc.Close()

			b, err := ioutil.ReadAll(rc)
			return string(b), err
		}
	}

	return "", fmt.Errorf("Could not find %s in %s", path, zipFile)
}
