package utils

import (
	"fmt"
	"github.com/wizard7414/epos_v2/domain"
	"github.com/wizard7414/epos_v2/domain/miner"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func SaveImageFromUrl(graphics miner.Graphics) (result bool, err error) {
	response, e := http.Get(graphics.Url)
	if e != nil {
		return false, e
	}
	defer response.Body.Close()

	file, err := os.Create(graphics.TargetUrl + graphics.TargetName + graphics.Extension)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SaveImageFromUrlWithVpn(graphics miner.Graphics) (result bool, err error) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	tbTransport := &http.Transport{Dial: dialer.Dial}

	request, requestErr := http.NewRequest("GET", graphics.Url, nil)
	if requestErr != nil {
		log.Fatalln(requestErr)
	}

	client := &http.Client{Transport: tbTransport}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Fatal(responseErr)
	}

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	file, err := os.Create(graphics.TargetUrl + graphics.TargetName + graphics.Extension)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetGraphicsUrl(object miner.ObjectV) string {
	for attrId := range object.Attributes {
		if object.Attributes[attrId].Code.Title == "graphics" {
			return object.Attributes[attrId].Value
		}
	}
	return ""
}

func createGraphicsByObject(object miner.ObjectV, saveUrl string) miner.Graphics {
	objectGraphics := miner.Graphics{
		Url:        GetGraphicsUrl(object),
		Extension:  "",
		TargetName: PrepareFileNameWithId(object.Url),
		TargetUrl:  saveUrl,
	}
	objectGraphics.Extension = GetExtension(objectGraphics)

	return objectGraphics
}

func saveGraph(objectGraphics miner.Graphics) bool {
	result, err := SaveImageFromUrl(objectGraphics)
	if err != nil {
		fmt.Println("==============================================================")
		fmt.Println(objectGraphics.TargetName + " - ERROR")
		fmt.Println("==============================================================")
		panic(err)
	}
	return result
}

func saveGraphVpn(objectGraphics miner.Graphics) bool {
	result, err := SaveImageFromUrlWithVpn(objectGraphics)
	if err != nil {
		fmt.Println("==============================================================")
		fmt.Println(objectGraphics.TargetName + " - ERROR")
		fmt.Println("==============================================================")
		panic(err)
	}
	return result
}

func SaveGraphicsStage(completesObjects []miner.ObjectV, config domain.EposConfig) {
	saveCount := 1

	for objectId := range completesObjects {
		objectGraphics := createGraphicsByObject(completesObjects[objectId], config.GraphicsPath)
		_, searchErr := os.Stat(objectGraphics.TargetUrl + objectGraphics.TargetName + objectGraphics.Extension)
		if searchErr != nil {
			if os.IsNotExist(searchErr) {
				var result bool
				if config.GraphVpnSave {
					result = saveGraphVpn(objectGraphics)
				} else {
					result = saveGraph(objectGraphics)
				}
				fmt.Println("[" + strconv.Itoa(saveCount) + "/" + strconv.Itoa(len(completesObjects)) + "] " + objectGraphics.TargetName + " - " + strconv.FormatBool(result))
			}
		} else {
			fmt.Println("[" + strconv.Itoa(saveCount) + "/" + strconv.Itoa(len(completesObjects)) + "] " + objectGraphics.TargetName + " -  skipped")
		}
		saveCount++
	}
}
