package helper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type IpfsUploadResposne struct {
	IpfsHash string
}

func UploadDataToIpfs(data []byte) (string, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	jsonBody := []byte(fmt.Sprintf(`{"code": "%s"}`, string(data)))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "https://explorer.litprotocol.com/api/pinata/upload", bodyReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(res.Body)

	// Decode the JSON data into a Person struct
	var ipfsUploadResposne IpfsUploadResposne
	if err := decoder.Decode(&ipfsUploadResposne); err != nil {
		return "", err
	}
	return ipfsUploadResposne.IpfsHash, nil
}
