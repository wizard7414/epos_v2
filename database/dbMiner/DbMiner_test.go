package dbMiner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"os"
	"testing"
	"time"
)

var base = DbMiner{
	Attribute:       AttrDao{nil},
	AttributeType:   AttrTypeDao{nil},
	AttributeCode:   AttrCodeDao{nil},
	Resource:        ResDao{nil},
	Source:          SrcDao{nil},
	Entry:           EntDao{nil},
	ObjectType:      ObjTypeDao{nil},
	Object:          ObjDao{nil},
	ObjectAttribute: ObjAttrDao{nil},
}

func TestMain(m *testing.M) {
	err := base.InitDb("../resource/epos-dbMiner.db")
	if err != nil {
		panic("Unable to start test set!")
	}
	defer base.CloseDb()

	exitRes := m.Run()

	os.Exit(exitRes)
}

func TestDbMiner_CreateAttribute(t *testing.T) {
	testAttribute := miner.AttrV{
		ID: 0,
		Code: miner.AttrCodeV{
			ID:    0,
			Title: "testCode",
		},
		AttrType: miner.AttrTypeV{
			ID:    0,
			Title: "testType",
		},
		Value: "testValue",
	}

	createErr := base.CreateAttribute(testAttribute)
	if createErr != nil {
		t.Fail()
	}
}

func TestDbMiner_CreateSource(t *testing.T) {
	testSource := miner.SourceV{
		ID:    0,
		Title: "Test Source",
		Url:   "Test source URL",
		Resource: miner.ResourceV{
			ID:    0,
			Title: "Test Resource",
			Url:   "Test resource URL",
		},
		ChangeDate: time.Now(),
	}

	createErr := base.CreateSource(testSource)
	if createErr != nil {
		t.Fail()
	}
}

func TestDbMiner_CreateEntry(t *testing.T) {
	testEntry := miner.EntryV{
		ID:    0,
		Title: "Test Entry",
		Url:   "Test entry Url",
		Source: miner.SourceV{
			ID:    0,
			Title: "Test Source",
			Url:   "Test source URL",
			Resource: miner.ResourceV{
				ID:    0,
				Title: "Test Resource",
				Url:   "Test resource URL",
			},
			ChangeDate: time.Now(),
		},
		AddDate: time.Time{},
	}

	createErr := base.CreateEntry(testEntry)
	if createErr != nil {
		t.Fail()
	}
}

func TestDbMiner_CreateObject(t *testing.T) {
	var attributes []miner.AttrV

	attributes = append(attributes, miner.AttrV{
		ID: 0,
		Code: miner.AttrCodeV{
			ID:    0,
			Title: "testCode",
		},
		AttrType: miner.AttrTypeV{
			ID:    0,
			Title: "testType",
		},
		Value: "testValue",
	})

	testObject := miner.ObjectV{
		ID:    0,
		Title: "Test Object",
		Entry: miner.EntryV{
			ID:    0,
			Title: "Test Entry",
			Url:   "Test entry Url",
			Source: miner.SourceV{
				ID:    0,
				Title: "Test Source",
				Url:   "Test source URL",
				Resource: miner.ResourceV{
					ID:    0,
					Title: "Test Resource",
					Url:   "Test resource URL",
				},
				ChangeDate: time.Now(),
			},
			AddDate: time.Time{},
		},
		Url:     "Test object Url",
		AddDate: time.Now(),
		ObjectType: miner.ObjTypeV{
			ID:    0,
			Title: "Test Object Type",
		},
		Attributes: attributes,
	}

	createErr := base.CreateObject(testObject)
	if createErr != nil {
		t.Fail()
	}
}

func TestEposDbSqlite_GetObjectAttributes(t *testing.T) {
	attributes, findError := base.GetObjectAttributes(1622233508865224101)
	if findError != nil || len(attributes) != 1 {
		t.Fail()
	}
}

func TestEposDbSqlite_GetObjectByUrl(t *testing.T) {
	objects, findErr := base.GetObjectByUrl("Test object Url")
	if findErr != nil || len(objects) != 1 {
		t.Fail()
	}
}
