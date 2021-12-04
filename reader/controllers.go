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
)

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
	if len(s) > 0 {
		return fmt.Sprintf("%s\n", s[len(s)-1])
	}
	return ""
}
