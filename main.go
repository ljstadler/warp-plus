package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numbers    = "0123456789"
)

var (
	ID               string
	ERROR_INTERVAL   int
	SUCCESS_INTERVAL int
	random           = rand.New(rand.NewSource(time.Now().UnixNano()))
	url              = fmt.Sprintf("https://api.cloudflareclient.com/v0a%s/reg", generateNumbers(3))
)

func generateCharacters(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[random.Intn(len(characters))]
	}
	return string(result)
}

func generateNumbers(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = numbers[random.Intn(len(numbers))]
	}
	return string(result)
}

func run() error {
	installId := generateCharacters(22)

	body, bodyBytes := map[string]any{
		"fcm_token":    installId + ":APA91b" + generateCharacters(134),
		"install_id":   installId,
		"key":          generateCharacters(43) + "=",
		"locale":       "en_US",
		"referrer":     ID,
		"tos":          time.Now().Format("2006-01-02T15:04:05.999-07:00"),
		"type":         "Android",
		"warp_enabled": false,
	}, new(bytes.Buffer)

	if err := json.NewEncoder(bodyBytes).Encode(body); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bodyBytes)
	if err != nil {
		return err
	}

	req.Header = map[string][]string{
		"Accept-Encoding": {"gzip"},
		"Connection":      {"Keep-Alive"},
		"Content-Type":    {"application/json; charset=UTF-8"},
		"Host":            {"api.cloudflareclient.com"},
		"User-Agent":      {"okhttp/3.12.1"},
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return nil
	} else {
		return errors.New(res.Status)
	}
}

func main() {
	ID = os.Getenv("ID")
	ERROR_INTERVAL, err := strconv.Atoi(os.Getenv("ERROR_INTERVAL"))
	if err != nil {
		log.Fatal(err)
	}
	SUCCESS_INTERVAL, err := strconv.Atoi(os.Getenv("SUCCESS_INTERVAL"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		err := run()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(ERROR_INTERVAL) * time.Second)
		} else {
			log.Println("1 GB Added")
			time.Sleep(time.Duration(SUCCESS_INTERVAL) * time.Second)
		}
	}
}
