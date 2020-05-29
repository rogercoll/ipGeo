package ipgeo

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	url = "http://api.ipstack.com/"
)

func makeUrl(ip, token string) (*http.Request, error) {
	tmpUrl := url + ip
	req, err := http.NewRequest("GET", tmpUrl, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("access_key", token)
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func GetInfo(ip, token string) IPStack {
	req, err := makeUrl(ip, token)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	info := IPStack{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Fatal(err)
	}
	return info
}
