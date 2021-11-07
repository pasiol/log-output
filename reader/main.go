package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func readUUIDs(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("opening file %s failed: %s", os.Getenv("APP_INPUT_FILE"), err)
	}
	scanner := bufio.NewScanner(f)
	var content []string
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}
	log.Printf("readed file %s", os.Getenv("APP_INPUT_FILE"))
	return content
}

func unwindLToS(s []string) string {
	o := ""
	for _, row := range s {
		o = o + row + "\n"
	}
	return o
}

func handler(w http.ResponseWriter, r *http.Request) {
	var uuids []string
	if fileExists(os.Getenv("APP_INPUT_FILE")) {
		uuids = readUUIDs(os.Getenv("APP_INPUT_FILE"))
	} else {
		log.Printf("opening input file failed: %s", os.Getenv("APP_INPUT_FILE"))
	}

	_, err := fmt.Fprintf(w, "%s", unwindLToS(uuids))
	if err != nil {
		log.Fatalf("writing response failed %s", r.RemoteAddr)
	}
}

func main() {

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8888"
	}

	if !fileExists(os.Getenv("APP_LOG_FILE")) {
		createFile(os.Getenv("APP_LOG_FILE"), "")
	}
	f, err := os.OpenFile(os.Getenv("APP_LOG_FILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Opening log file failed: %s", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("closing file failed %s %s", os.Getenv("APP_LOG_FILE"), err)
		}
	}(f)
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Printf("starting reader")

	address := "0.0.0.0:" + os.Getenv("APP_PORT")
	log.Printf("starting in address %s.", address)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
