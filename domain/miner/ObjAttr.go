package miner

import "time"

type ObjAttrS struct {
	Object    int64
	Attribute int64
}

func (ObjAttrS) TableName() string {
	return "object_attribute"
}

func (store ObjAttrS) TransformToView() ObjAttrV {
	return ObjAttrV{
		Object: ObjectV{
			ID:         store.Object,
			Title:      "",
			Entry:      EntryV{},
			Url:        "",
			AddDate:    time.Now(),
			ObjectType: ObjTypeV{},
			Attributes: []AttrV{},
		},
		Attribute: AttrV{
			ID:       store.Attribute,
			Code:     AttrCodeV{},
			AttrType: AttrTypeV{},
			Value:    "",
		},
	}
}

type ObjAttrV struct {
	Object    ObjectV
	Attribute AttrV
}

func (view ObjAttrV) TransformToStore() ObjAttrS {
	return ObjAttrS{
		Object:    view.Object.ID,
		Attribute: view.Attribute.ID,
	}
}
