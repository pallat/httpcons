package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: time.Second * 1,
	}

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("call successfully")
}
