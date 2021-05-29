package utils

import (
	"fmt"
	"github.com/wizard7414/epos_v2/domain/miner"
	"strings"
	"testing"
)

func TestPrepareFileName(t *testing.T) {
	str := PrepareFileName("DDD sdfsdf: sdfsd \\ efwefsd ? / sdfsdfesfdsG")
	fmt.Println(str)
	if strings.Contains(str, "\\") {
		t.Fail()
	}
}

func TestSetExtension(t *testing.T) {
	testGraphics := miner.Graphics{
		Url:        "http://test/path/image.png",
		Extension:  "",
		TargetName: "newName",
		TargetUrl:  "/home/temp",
	}

	extension := GetExtension(testGraphics)
	fmt.Println(extension)
	if extension != ".png" {
		t.Fail()
	}
}
