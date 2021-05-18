package miner

type ResourceS struct {
	ID    int64
	Title string
	Url   string
}

func (ResourceS) TableName() string {
	return "resource"
}

func (store ResourceS) TransformToView() ResourceV {
	return ResourceV{
		ID:    store.ID,
		Title: store.Title,
		Url:   store.Url,
	}
}

type ResourceV struct {
	ID    int64
	Title string
	Url   string
}

func (view ResourceV) TransformToStore() ResourceS {
	return ResourceS{
		ID:    view.ID,
		Title: view.Title,
		Url:   view.Url,
	}
}
