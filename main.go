package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	for {
		u := uuid.New()
		localTime := time.Now()
		fmt.Printf("%v %s\n", localTime.UTC().Format(time.RFC3339Nano), u.String())
		time.Sleep(5 * time.Second)
	}
}
