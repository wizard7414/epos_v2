package miner

import (
	"github.com/wizard7414/epos_v2/domain/miner"
	"testing"
)

func TestResDao_Create(t *testing.T) {
	resource := miner.ResourceS{
		ID:    0,
		Title: "testResource",
		Url:   "testResourceUrl",
	}

	createErr := base.Resource.Create(resource)
	if createErr != nil {
		t.Fail()
	}
}

func TestResDao_GetById(t *testing.T) {
	resource, findErr := base.Resource.GetById(1)

	if resource.Title != "testResource" || findErr != nil {
		t.Fail()
	}
}

func TestResDao_GetByTitle(t *testing.T) {
	resourceList, findErr := base.Resource.GetByTitle("testResource")

	if len(resourceList) != 1 || resourceList[0].Title != "testResource" || findErr != nil {
		t.Fail()
	}
}

func TestResDao_Delete(t *testing.T) {
	result := base.Resource.Delete(1)

	if result != nil {
		t.Fail()
	}
}
