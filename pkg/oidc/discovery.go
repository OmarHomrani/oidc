package oidc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseDiscovery(url string) (Discovery, error) {
	fmt.Println("ParseDiscovery%s", url)
	var discovery Discovery
	res, err := http.Get(url)
	if err != nil {
		return discovery, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return discovery, err
	}
	if err = json.Unmarshal(body, &discovery); err != nil {
		return discovery, err
	}
	return discovery, nil
}
