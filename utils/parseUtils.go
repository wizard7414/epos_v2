package utils

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
)

func GetHtmlDocFormUrl(url string) (*goquery.Document, int) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		if response.StatusCode == 403 {
			log.Printf("status code error: %d %s", response.StatusCode, response.Status)
			return nil, response.StatusCode
		} else {
			log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
		}
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, 200
}

func GetHtmlDocFromUrlWithAuth(pathUrl string, authHeader string) (*goquery.Document, int) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	tbTransport := &http.Transport{Dial: dialer.Dial}

	request, requestErr := http.NewRequest("GET", pathUrl, nil)
	if requestErr != nil {
		log.Fatalln(requestErr)
	}

	request.Header.Add("Cookie", authHeader)

	client := &http.Client{Transport: tbTransport}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Fatal(responseErr)
	}

	if response.StatusCode != 200 {
		if response.StatusCode == 403 {
			log.Printf("status code error: %d %s", response.StatusCode, response.Status)
			return nil, response.StatusCode
		} else {
			log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
		}
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, 200
}
