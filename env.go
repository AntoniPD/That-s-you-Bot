package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func initEnv() {
	data, err := readLines(".env")
	check(err)
	os.Setenv("PAGE_ACCESS_TOKEN", data[0])
	os.Setenv("VERIFY_TOKEN", data[1])
	os.Setenv("POST_MESSAGE", "https://graph.facebook.com/v2.6/me/messages?access_token="+os.Getenv("PAGE_ACCESS_TOKEN"))
	os.Setenv("GET_USER_INFO", "https://graph.facebook.com/v2.6/%s?access_token="+os.Getenv("PAGE_ACCESS_TOKEN"))
}
