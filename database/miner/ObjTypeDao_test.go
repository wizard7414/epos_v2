package miner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestObjTypeDao_Create(t *testing.T) {
	objectType := miner.ObjTypeS{
		ID:    0,
		Title: "testObjectType",
	}

	createErr := base.ObjectType.Create(objectType)
	if createErr != nil {
		t.Fail()
	}
}

func TestObjTypeDao_GetById(t *testing.T) {
	objectType, findErr := base.ObjectType.GetById(1)

	if objectType.Title != "testObjectType" || findErr != nil {
		t.Fail()
	}
}

func TestObjTypeDao_GetByTitle(t *testing.T) {
	objectTypes, findErr := base.ObjectType.GetByTitle("testObjectType")

	if len(objectTypes) != 1 || objectTypes[0].Title != "testObjectType" || findErr != nil {
		t.Fail()
	}
}

func TestObjTypeDao_Delete(t *testing.T) {
	result := base.ObjectType.Delete(1)

	if result != nil {
		t.Fail()
	}
}
