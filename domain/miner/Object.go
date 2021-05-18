package miner

import "time"

type ObjectS struct {
	ID         int64
	Title      string
	Entry      int64
	Url        string
	AddDate    int64
	ObjectType int64
}

func (ObjectS) TableName() string {
	return "object"
}

func (store ObjectS) TransformToView() ObjectV {
	return ObjectV{
		ID:    store.ID,
		Title: store.Title,
		Entry: EntryV{
			ID:      store.Entry,
			Title:   "",
			Url:     "",
			Source:  SourceV{},
			AddDate: time.Now(),
		},
		Url:     store.Url,
		AddDate: time.Unix(store.AddDate, 0),
		ObjectType: ObjTypeV{
			ID:    store.ObjectType,
			Title: "",
		},
		Attributes: []AttrV{},
	}
}

type ObjectV struct {
	ID         int64
	Title      string
	Entry      EntryV
	Url        string
	AddDate    time.Time
	ObjectType ObjTypeV
	Attributes []AttrV
}

func (view ObjectV) TransformToStore() ObjectS {
	return ObjectS{
		ID:         view.ID,
		Title:      view.Title,
		Entry:      view.Entry.ID,
		Url:        view.Url,
		AddDate:    view.AddDate.Unix(),
		ObjectType: view.ObjectType.ID,
	}
}
