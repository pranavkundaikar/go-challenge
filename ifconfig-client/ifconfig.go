package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch_ip(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {

	url := "http://ifconfig.me/ip"

	fmt.Println("Trying to fetch public from ifconfig.me")

	ip := fetch_ip(url)

	fmt.Println("This is the IP address: ", ip)

}
