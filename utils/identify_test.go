package utils

import (
	"fmt"
	"testing"
)

func TestGetTimestamp(t *testing.T) {
	timestampId := GetTimestamp()

	fmt.Println(timestampId)

	if timestampId <= 0 {
		t.Fail()
	}
}
