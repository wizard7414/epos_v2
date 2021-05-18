package miner

import "time"

type EntryS struct {
	ID      int64
	Title   string
	Url     string
	Source  int64
	AddDate int64
}

func (EntryS) TableName() string {
	return "entry"
}

func (store EntryS) TransformToView() EntryV {
	return EntryV{
		ID:    store.ID,
		Title: store.Title,
		Url:   store.Url,
		Source: SourceV{
			ID:         store.Source,
			Title:      "",
			Url:        "",
			Resource:   ResourceV{},
			ChangeDate: time.Now(),
		},
		AddDate: time.Unix(store.AddDate, 0),
	}
}

type EntryV struct {
	ID      int64
	Title   string
	Url     string
	Source  SourceV
	AddDate time.Time
}

func (view EntryV) TransformToStore() EntryS {
	return EntryS{
		ID:      view.ID,
		Title:   view.Title,
		Url:     view.Url,
		Source:  view.Source.ID,
		AddDate: view.AddDate.Unix(),
	}
}
