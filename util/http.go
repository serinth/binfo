package util

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {

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
