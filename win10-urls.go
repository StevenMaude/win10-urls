package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/xmlpath.v2"
)

// downloadHTML downloads the Windows 10 download page HTML.
func downloadHTML() (string, error) {
	fmt.Println("Retrieving data from microsoft.com...")
	url := "https://www.microsoft.com/en-us/api/controls/contentinclude/html?pageId=cfa9e580-a81e-4a4b-a846-7b21bf4e2e5b&host=www.microsoft.com&segments=software-download,windows10ISO&query=&action=GetProductDownloadLinksBySku&skuId=3864&language=English&sdVersion=2"

	resp, err := http.Post(url, "", strings.NewReader(""))
	if err != nil {
		fmt.Println("error retrieving data.")
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response.")
		return "", err
	}

	return string(body), nil
}

// parseHTML parses the Windows 10 download page HTML.
func parseHTML(body string) error {
	r := strings.NewReader(body)
	root, err := xmlpath.ParseHTML(r)
	if err != nil {
		fmt.Println("error parsing HTML.")
		return err
	}

	path := xmlpath.MustCompile("//div[@class]/a/@href")
	iter := path.Iter(root)

	for iter.Next() {
		fmt.Println(iter.Node().String())
	}
	return nil
}

// main retrieves Windows 10 download URLs.
func main() {
	body, err := downloadHTML()
	if err != nil {
		os.Exit(1)
	}

	err = parseHTML(body)
	if err != nil {
		os.Exit(1)
	}
}
