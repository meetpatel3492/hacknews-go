package clients

import (
	"encoding/json"
	"net/http"
	"time"
)

var HackerNewsHttpClient = &http.Client{Timeout: 5 * time.Second}

func GetStructFromJson(url string, target interface{}) error {
	response, err := HackerNewsHttpClient.Get(url)
	if err != nil{
		return err
	}
	defer response.Body.Close()
	return structFromJsonResponse(response, target)
}

func structFromJsonResponse(resp *http.Response,target interface{}) error {
	return json.NewDecoder(resp.Body).Decode(target)
}