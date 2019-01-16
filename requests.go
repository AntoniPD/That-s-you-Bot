package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"msgbot/models/messenger"
	"net/http"
	"os"
)

// PostMessage -
// Executing a post request which sends the message to the user.
func PostMessage(url string, body io.Reader) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

// GetUserInfo -
// Gets {firstName, lastName, profilePic, id} of a user
// which already had an interaction with the bot.
func GetUserInfo(userId string) models.UserInfo {
	url := fmt.Sprintf(os.Getenv("GET_USER_INFO"), userId)
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := models.UserInfo{}
	jsonErr := json.Unmarshal(body, &user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	defer res.Body.Close()
	return user
}
