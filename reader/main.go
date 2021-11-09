package main

import (
	"bufio"
	"errors"
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

func readPingPongService() (string, error) {
	url := fmt.Sprintf("http://%s/pingpong", os.Getenv("APP_PING_PONG_HOST"))
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		log.Printf("get request status code: %d", response.StatusCode)
		return "", errors.New("get request failed")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("closing get request body failed")
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func lastRow(s []string) string {
	return fmt.Sprintf("%s\n", s[len(s)-1])
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("getting request from %s", r.RemoteAddr)
	var uuids []string
	if fileExists(os.Getenv("APP_INPUT_FILE")) {
		uuids = readUUIDs(os.Getenv("APP_INPUT_FILE"))
	} else {
		log.Printf("opening input file failed: %s", os.Getenv("APP_INPUT_FILE"))
	}

	pingpong, err := readPingPongService()
	if err != nil {
		log.Printf("reading pingpong service failed: %s", err)
	}

	bytes, err := fmt.Fprintf(w, "%s%s", lastRow(uuids), pingpong)
	if err != nil {
		log.Fatalf("writing response failed %s", r.RemoteAddr)
	}
	log.Printf("%d bytes written %s", bytes, r.RemoteAddr)
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

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Printf("starting reader")

	log.Printf("starting in address 0.0.0.0:%s.", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
