package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func createDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("creating directory failed: %s", err)
	}
}

func createFile(filename string, content string) {
	if !fileExists(filepath.Dir(filename)) {
		createDir(filepath.Dir(filename))
	}
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("creating file %s failed: %s", filename, err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
