package miner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestAttrDao_Create(t *testing.T) {
	attribute := miner.AttrS{
		ID:            0,
		Code:          0,
		AttributeType: 0,
		Value:         "testValue",
	}

	createErr := base.Attribute.Create(attribute)
	if createErr != nil {
		t.Fail()
	}
}

func TestAttrDao_GetById(t *testing.T) {
	attribute, findErr := base.Attribute.GetById(1)

	if attribute.Value != "testValue" || findErr != nil {
		t.Fail()
	}
}

func TestAttrDao_GetByCode(t *testing.T) {
	attributeList, findErr := base.Attribute.GetByCode(0)

	if len(attributeList) != 1 || attributeList[0].Value != "testValue" || findErr != nil {
		t.Fail()
	}
}

func TestAttrDao_Delete(t *testing.T) {
	result := base.Attribute.Delete(1)

	if result != nil {
		t.Fail()
	}
}
