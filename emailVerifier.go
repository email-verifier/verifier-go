package emailVerifier

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func verify(email, access_token string) (string, string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://verifier.meetchopra.com/verify/", email, "?token=", access_token, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
  return bodyText
}
