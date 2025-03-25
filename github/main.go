package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if _, _, err := githubInfo("sramanathan"); err != nil {
		fmt.Printf("Issue with gathering info: %s", err)
	}
}

type Reply struct {
	Name         string `json:"name,omitempty"`
	Public_Repos int    `json:"public_repos,omitempty"`
}

func githubInfo(login string) (string, int, error) {
	fmt.Printf("Login: %s", login)

	resp1, err := callMockAPI()
	if err != nil {
		log.Fatalf("error: can't decode - %s", err)
		return "", 0, err
	}

	resp2, err := callMockAPI()
	if err != nil {
		log.Fatalf("error: can't decode - %s", err)
		return "", 0, err
	}

	var r Reply
	var anonymousR struct {
		Name         string `json:"name,omitempty"`
		Public_Repos int    `json:"public_repos,omitempty"`
	}

	dec1 := json.NewDecoder(resp1.Body)
	if err := dec1.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
		return "", 0, err
	}

	dec2 := json.NewDecoder(resp2.Body)
	if err := dec2.Decode(&anonymousR); err != nil {
		log.Fatalf("error: can't decode - %s", err)
		return "", 0, err
	}

	fmt.Printf("%v\n", r)
	fmt.Printf("%v\n", anonymousR)

	return r.Name, r.Public_Repos, err
}

func callMockAPI() (resp *http.Response, err error) {
	url := "http://localhost:3000/content"

	resp, err = http.Post(url, "application/json", nil)
	if err != nil {
		log.Fatalf("error: %s", err)
		return resp, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
		return resp, err
	}

	fmt.Printf("Content-Type: %s \n", resp.Header.Get("Content-Type"))

	return resp, err
}
