package dbMiner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestAttrTypeDao_Create(t *testing.T) {
	attributeType := miner.AttrTypeS{
		ID:    0,
		Title: "testType",
	}

	createErr := base.AttributeType.Create(attributeType)
	if createErr != nil {
		t.Fail()
	}
}

func TestAttrTypeDao_GetById(t *testing.T) {
	attributeType, findErr := base.AttributeType.GetById(1)

	if attributeType.Title != "testType" || findErr != nil {
		t.Fail()
	}
}

func TestAttrTypeDao_GetByTitle(t *testing.T) {
	attributeTypeList, findErr := base.AttributeType.GetByTitle("testType")

	if len(attributeTypeList) != 1 || attributeTypeList[0].Title != "testType" || findErr != nil {
		t.Fail()
	}
}

func TestAttrTypeDao_Delete(t *testing.T) {
	result := base.AttributeType.Delete(1)

	if result != nil {
		t.Fail()
	}
}
