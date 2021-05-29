package utils

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestSaveImageFromUrl(t *testing.T) {
	testGraphics := miner.Graphics{
		Url:        "https://yastatic.net/iconostasis/_/5mdPq4V7ghRgzBvMkCaTzd2fjYg.png",
		Extension:  ".png",
		TargetName: "test",
		TargetUrl:  "/home/wizard/Загрузки/",
	}

	result, _ := SaveImageFromUrl(testGraphics)
	if result != true {
		t.Fail()
	}
}

func TestSaveImageFromUrlFail(t *testing.T) {
	testGraphics := miner.Graphics{
		Url:        "",
		Extension:  ".png",
		TargetName: "test",
		TargetUrl:  "/home/wizard/Загрузки/",
	}

	_, openErr := SaveImageFromUrl(testGraphics)
	if openErr == nil {
		t.Fail()
	}

	testGraphics.Url = "https://yastatic.net/iconostasis/_/5mdPq4V7ghRgzBvMkCaTzd2fjYg.png"
	testGraphics.TargetUrl = "/nil"

	_, saveErr := SaveImageFromUrl(testGraphics)
	if saveErr == nil {
		t.Fail()
	}
}
