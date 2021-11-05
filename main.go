package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request %s, %s", r.Method, r.URL)
	u := uuid.New()
	localTime := time.Now()
	fmt.Printf("%v %s\n", localTime.UTC().Format(time.RFC3339Nano), u.String())
	bytes, err := fmt.Fprintf(w, "%s %s", localTime.UTC().Format(time.RFC3339Nano), u.String())
	if err != nil {
		log.Fatal("writing the response failed")
	}
	log.Printf("%d bytes written", bytes)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf(" Server started in port %s.", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
