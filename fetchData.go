package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchData(url string, result any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(result)
}
