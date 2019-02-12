package emailVerifier

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Verfier_struct struct {
  Status bool `json:"status"`
}

func verify(email, access_token string) bool {
	var url = "https://verifier.meetchopra.com/verify/"+ email+ "?token="+ access_token
	client := &http.Client{}
	req, err := http.NewRequest("GET", url , nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(resp.Body)
	var data Verfier_struct
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data.Status
}
