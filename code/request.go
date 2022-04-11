package code

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func MakeRequest(myUrl string) *goquery.Document {

	fmt.Println("URL:>", myUrl)

	form := url.Values{}


	req, err := http.NewRequest("GET", myUrl, strings.NewReader(form.Encode()))
	req.Header.Set("Accept-Language", "ru")
	req.Header.Set("User-Agent", GetAgents())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}