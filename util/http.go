package util

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	f, err := os.OpenFile("binfo_http_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer f.Close()

	log.SetOutput(f)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	if err != nil {
		log.Println(err)
		return err
	}

	r, err := client.Do(req)

	if err != nil {
		log.Println("Error fetching from URL")
		return err
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
