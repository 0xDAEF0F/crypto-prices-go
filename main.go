package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	hl_url := "https://api.hyperliquid.xyz"

	data := []byte(`{"type": "allMids"}`)
	resp, err := http.Post(hl_url+"/info", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Printf("%+v\n", result)
}
