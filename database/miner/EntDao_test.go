package miner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
	"time"
)

func TestEntDao_Create(t *testing.T) {
	entry := miner.EntryS{
		ID:      0,
		Title:   "testEntry",
		Url:     "testEntryUrl",
		Source:  0,
		AddDate: time.Now().Unix(),
	}

	createErr := base.Entry.Create(entry)
	if createErr != nil {
		t.Fail()
	}
}

func TestEntDao_GetById(t *testing.T) {
	entry, findErr := base.Entry.GetById(1)

	if entry.Title != "testEntry" || findErr != nil {
		t.Fail()
	}
}

func TestEntDao_GetByTitle(t *testing.T) {
	entries, findErr := base.Entry.GetByTitle("testEntry")

	if len(entries) != 1 || entries[0].Title != "testEntry" || findErr != nil {
		t.Fail()
	}
}

func TestEntDao_Delete(t *testing.T) {
	result := base.Entry.Delete(1)

	if result != nil {
		t.Fail()
	}
}
