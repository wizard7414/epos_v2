package dbMiner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestAttrCodeDao_Create(t *testing.T) {
	attributeCode := miner.AttrCodeS{
		ID:    0,
		Title: "testCode",
	}

	createErr := base.AttributeCode.Create(attributeCode)
	if createErr != nil {
		t.Fail()
	}
}

func TestAttrCodeDao_GetById(t *testing.T) {
	attributeCode, findErr := base.AttributeCode.GetById(1)
	if findErr != nil || attributeCode.Title != "testCode" {
		t.Fail()
	}
}

func TestAttrCodeDao_GetByTitle(t *testing.T) {
	attributeCodeList, findErr := base.AttributeCode.GetByTitle("testCode")

	if len(attributeCodeList) != 1 || attributeCodeList[0].Title != "testCode" || findErr != nil {
		t.Fail()
	}
}

func TestAttrCodeDao_Delete(t *testing.T) {
	result := base.AttributeCode.Delete(1)

	if result != nil {
		t.Fail()
	}
}
