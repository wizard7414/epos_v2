package miner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
	"time"
)

func TestObjDao_Create(t *testing.T) {
	object := miner.ObjectS{
		ID:         1,
		Title:      "testObject",
		Entry:      1,
		Url:        "testObjectUrl",
		AddDate:    time.Now().Unix(),
		ObjectType: 1,
	}

	createErr := base.Object.Create(object)
	if createErr != nil {
		t.Fail()
	}
}

func TestObjDao_GetById(t *testing.T) {
	object, findErr := base.Object.GetById(1)

	if object.Title != "testObject" || findErr != nil {
		t.Fail()
	}
}

func TestObjDao_GetByTitle(t *testing.T) {
	objects, findErr := base.Object.GetByTitle("testObject")

	if len(objects) != 1 || objects[0].Title != "testObject" || findErr != nil {
		t.Fail()
	}
}

func TestObjDao_GetByUrl(t *testing.T) {
	objects, findErr := base.Object.GetByUrl("testObjectUrl")

	if len(objects) != 1 || objects[0].Title != "testObject" || findErr != nil {
		t.Fail()
	}
}

func TestObjDao_Delete(t *testing.T) {
	result := base.Object.Delete(1)

	if result != nil {
		t.Fail()
	}
}
