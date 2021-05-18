package miner

import "time"

type SourceS struct {
	ID         int64
	Title      string
	Url        string
	Resource   int64
	ChangeDate int64
}

func (SourceS) TableName() string {
	return "source"
}

func (store SourceS) TransformToView() SourceV {
	return SourceV{
		ID:    store.ID,
		Title: store.Title,
		Url:   store.Url,
		Resource: ResourceV{
			ID:    store.Resource,
			Title: "",
			Url:   "",
		},
		ChangeDate: time.Unix(store.ChangeDate, 0),
	}
}

type SourceV struct {
	ID         int64
	Title      string
	Url        string
	Resource   ResourceV
	ChangeDate time.Time
}

func (view SourceV) TransformToStore() SourceS {
	return SourceS{
		ID:         view.ID,
		Title:      view.Title,
		Url:        view.Url,
		Resource:   view.Resource.ID,
		ChangeDate: view.ChangeDate.Unix(),
	}
}
