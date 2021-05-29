package deviant

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wizard7414/epos_v2/database/dbMiner"
	"github.com/wizard7414/epos_v2/domain/miner"
	"github.com/wizard7414/epos_v2/utils"
	"os"
	"strconv"
	"time"
)

func fillObject(doc *goquery.Document, path string) miner.ObjectV {
	var item miner.ObjectV

	item.ID = 0
	item.Title = FindObjectTitle(doc)
	item.Entry = FindObjectEntry(doc)
	item.Attributes = FindObjectAttributes(doc)
	item.AddDate = time.Now()
	item.Url = path
	item.ObjectType = miner.ObjTypeV{
		ID:    0,
		Title: "artwork",
	}

	return item
}

func ParseFromUrl(pathUrl string) (miner.ObjectV, int) {
	document, result := utils.GetHtmlDocFormUrl(pathUrl)
	if result == 200 {
		return fillObject(document, pathUrl), result
	}

	return miner.ObjectV{}, result
}

func ParseFromUrlWithAuth(pathUrl string, authHeader string) (miner.ObjectV, int) {
	document, result := utils.GetHtmlDocFromUrlWithAuth(pathUrl, authHeader)
	if result == 200 {
		return fillObject(document, pathUrl), result
	}

	return miner.ObjectV{}, result
}

func getGraphicsUrl(object miner.ObjectV) string {
	for attrId := range object.Attributes {
		if object.Attributes[attrId].Code.Title == "graphics" {
			return object.Attributes[attrId].Value
		}
	}
	return ""
}

func parseFromDeviant(path string) (miner.ObjectV, int, bool) {
	parsedObject, parseResult := ParseFromUrl(path)
	if parseResult == 200 {
		validateResult := ValidateParsedObject(parsedObject)
		return parsedObject, 200, validateResult
	}
	return parsedObject, parseResult, false
}

func parseFromDeviantWithAuth(path string, authHeader string) (miner.ObjectV, int, bool) {
	parsedObject, parseResult := ParseFromUrlWithAuth(path, authHeader)
	if parseResult == 200 {
		validateResult := ValidateParsedObject(parsedObject)
		return parsedObject, 200, validateResult
	}
	return parsedObject, parseResult, false
}

func ParseStage(base *dbMiner.DbMiner, bufferPath string, authHeader string) ([]miner.ObjectV, []miner.ObjectV) {
	var completesObjects []miner.ObjectV
	var rejectedObjects []miner.ObjectV

	strList, err := utils.ParseLinksFromFile(bufferPath)
	parseCount := len(strList)
	if err == nil {
		for id := range strList {
			echo := utils.EchoCount(parseCount, id)
			storedObjects, findErr := base.GetObjectByUrl(strList[id])
			if findErr == nil && len(storedObjects) == 0 {
				var parsedObject miner.ObjectV
				var parseResult int
				var validateResult bool
				if authHeader != "" {
					parsedObject, parseResult, validateResult = parseFromDeviantWithAuth(strList[id], authHeader)
				} else {
					parsedObject, parseResult, validateResult = parseFromDeviant(strList[id])
				}

				if parseResult == 403 {
					fmt.Println("==============================================================")
					fmt.Println(echo + "Stopped parse on object: " + strList[id])
					fmt.Println("==============================================================")
					break
				}

				if validateResult {
					completesObjects = append(completesObjects, parsedObject)
					fmt.Println(echo + "Complete parse object: " + strList[id])
				} else {
					rejectedObjects = append(rejectedObjects, parsedObject)
					fmt.Println(echo + "Reject parse object: " + strList[id])
				}
			} else {
				fmt.Println(echo + "Skipped stored object: " + strList[id])
			}
		}
	}

	fmt.Println("==============================================================")
	fmt.Println("Count of parsed objects: " + strconv.Itoa(len(completesObjects)))
	fmt.Println("Count of rejected objects: " + strconv.Itoa(len(rejectedObjects)))
	fmt.Println("==============================================================")

	return completesObjects, rejectedObjects
}

func SaveObjectsStage(base *dbMiner.DbMiner, completesObjects []miner.ObjectV) {
	parseCount := len(completesObjects)
	for objectId := range completesObjects {
		echo := utils.EchoCount(parseCount, objectId)
		object := completesObjects[objectId]
		createErr := base.CreateObject(object)
		if createErr != nil {
			panic("Create object Error!")
		}
		fmt.Println(echo + object.Title + " - saved")
	}
}

func createGraphicsByObject(object miner.ObjectV, saveUrl string) miner.Graphics {
	objectGraphics := miner.Graphics{
		Url:        getGraphicsUrl(object),
		Extension:  "",
		TargetName: utils.PrepareFileNameWithId(object.Url),
		TargetUrl:  saveUrl,
	}
	objectGraphics.Extension = utils.GetExtension(objectGraphics)

	return objectGraphics
}

func SaveGraphStage(completesObjects []miner.ObjectV, saveUrl string) {
	parseCount := len(completesObjects)
	for objectId := range completesObjects {
		echo := utils.EchoCount(parseCount, objectId)
		objectGraphics := createGraphicsByObject(completesObjects[objectId], saveUrl)

		_, searchErr := os.Stat(objectGraphics.TargetUrl + objectGraphics.TargetName + objectGraphics.Extension)
		if searchErr != nil {
			if os.IsNotExist(searchErr) {
				result, err := utils.SaveImageFromUrl(objectGraphics)
				if err != nil {
					panic(err)
				}
				fmt.Println(echo + objectGraphics.TargetName + " - " + strconv.FormatBool(result))
			}
		} else {
			fmt.Println(echo + objectGraphics.TargetName + " -  skipped")
		}
	}
}

func SaveGraphStageVpn(completesObjects []miner.ObjectV, saveUrl string) {
	for objectId := range completesObjects {
		objectGraphics := createGraphicsByObject(completesObjects[objectId], saveUrl)

		_, searchErr := os.Stat(objectGraphics.TargetUrl + objectGraphics.TargetName + objectGraphics.Extension)
		if searchErr != nil {
			if os.IsNotExist(searchErr) {
				result, err := utils.SaveImageFromUrlWithVpn(objectGraphics)
				if err != nil {
					fmt.Println("==============================================================")
					fmt.Println(completesObjects[objectId].Url + " - ERROR")
					fmt.Println("==============================================================")
					panic(err)
				}
				fmt.Println(objectGraphics.TargetName + " - " + strconv.FormatBool(result))
			}
		} else {
			fmt.Println(objectGraphics.TargetName + " -  skipped")
		}
	}
}
