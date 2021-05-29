package dbMiner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestObjAttrDao_Create(t *testing.T) {
	objectAttribute := miner.ObjAttrS{
		Object:    0,
		Attribute: 0,
	}

	createErr := base.ObjectAttribute.Create(objectAttribute)
	if createErr != nil {
		t.Fail()
	}
}

func TestObjAttrDao_GetById(t *testing.T) {
	objectAttribute, findErr := base.ObjectAttribute.GetById(0, 0)

	if objectAttribute.Attribute != 0 || objectAttribute.Object != 0 || findErr != nil {
		t.Fail()
	}
}

func TestObjAttrDao_GetByObject(t *testing.T) {
	objectAttributes, findErr := base.ObjectAttribute.GetByObject(0)

	if len(objectAttributes) != 1 || objectAttributes[0].Object != 0 || findErr != nil {
		t.Fail()
	}
}

func TestObjAttrDao_GetByAttribute(t *testing.T) {
	objectAttributes, findErr := base.ObjectAttribute.GetByAttribute(0)

	if len(objectAttributes) != 1 || objectAttributes[0].Attribute != 0 || findErr != nil {
		t.Fail()
	}
}

func TestObjAttrDao_Delete(t *testing.T) {
	result := base.ObjectAttribute.Delete(0, 0)

	if result != nil {
		t.Fail()
	}
}
