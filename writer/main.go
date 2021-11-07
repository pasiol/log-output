package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
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

func appendString2File(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("opening output file failed %s %s", filename, err)
	}
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatalf("appending string to file %s %s", filename, err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("closing file failed %s %s", filename, err)
	}

}

func writeUUID() {
	if !fileExists(os.Getenv("APP_OUTPUT_FILE")) {
		createFile(os.Getenv("APP_OUTPUT_FILE"), "")
	}
	for true {
		u := uuid.New()
		localTime := time.Now()

		o := fmt.Sprintf("%v %s\n", localTime.UTC().Format(time.RFC3339Nano), u.String())
		appendString2File(os.Getenv("APP_OUTPUT_FILE"), o)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	if !fileExists(os.Getenv("APP_LOG_FILE")) {
		createFile(os.Getenv("APP_LOG_FILE"), "")
	}
	f, err := os.OpenFile(os.Getenv("APP_LOG_FILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("opening log file failed: %s", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("closing file failed %s %s", os.Getenv("APP_LOG_FILE"), err)
		}
	}(f)
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Printf("starting writer")
	writeUUID()
}
