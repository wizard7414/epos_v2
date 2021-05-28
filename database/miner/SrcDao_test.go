package miner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
	"time"
)

func TestSrcDao_Create(t *testing.T) {
	source := miner.SourceS{
		ID:         0,
		Title:      "testSource",
		Url:        "testSourceUrl",
		Resource:   0,
		ChangeDate: time.Now().Unix(),
	}

	createErr := base.Source.Create(source)
	if createErr != nil {
		t.Fail()
	}
}

func TestSrcDao_GetById(t *testing.T) {
	source, findErr := base.Source.GetById(1)

	if source.Title != "testSource" || findErr != nil {
		t.Fail()
	}
}

func TestSrcDao_GetByTitle(t *testing.T) {
	sourceList, findErr := base.Source.GetByTitle("testSource")

	if len(sourceList) != 1 || sourceList[0].Title != "testSource" || findErr != nil {
		t.Fail()
	}
}

func TestSrcDao_UpdateDateById(t *testing.T) {
	updateErr := base.Source.UpdateDateById(1, time.Now().Unix())

	if updateErr != nil {
		t.Fail()
	}
}

func TestSrcDao_UpdateDateByTitle(t *testing.T) {
	updateErr := base.Source.UpdateDateByTitle("testSource", time.Now().Unix())

	if updateErr != nil {
		t.Fail()
	}
}

func TestSrcDao_Delete(t *testing.T) {
	result := base.Source.Delete(1)

	if result != nil {
		t.Fail()
	}
}
