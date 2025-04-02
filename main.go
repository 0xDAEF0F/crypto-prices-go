package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	hl_url := "https://api.hyperliquid.xyz"

	data := []byte(`{"type": "allMids"}`)
	resp, err := http.Post(hl_url+"/info", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error sending http request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading from body:", err)
		return
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for key, value := range result {
		midPrice, _ := strconv.ParseFloat(value, 64)
		fmt.Printf("%s: %f\n", key, midPrice)
	}
}
