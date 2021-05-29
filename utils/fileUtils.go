package utils

import (
	"bufio"
	"github.com/wizard7414/epos_v2/domain/miner"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFiles(baseFolder string) ([]os.FileInfo, error) {
	var files []os.FileInfo

	err := filepath.Walk(baseFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info)
		}
		return nil
	})

	return files, err
}

func PrepareFileName(fileName string) string {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	return re.ReplaceAllString(fileName, "")
}

func PrepareFileNameWithId(urlString string) string {
	strArr := strings.Split(urlString, "/")
	return GetGraphicsId(strArr[len(strArr)-1])
}

func GetGraphicsId(urlString string) string {
	strArr := strings.Split(urlString, "-")
	return strArr[len(strArr)-1]
}

func ParseLinksFromFile(pathFile string) ([]string, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func GetExtension(graphics miner.Graphics) string {
	validExtensions := []string{".jpg", ".JPG", ".jpeg", ".webp", ".png", ".gif"}
	var extension = ""

	for id := range validExtensions {
		if strings.Contains(graphics.Url, validExtensions[id]) {
			extension = validExtensions[id]
		}
	}

	return extension
}
