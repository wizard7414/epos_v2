package miner

type AttrCodeS struct {
	ID    int64
	Title string
}

func (AttrCodeS) TableName() string {
	return "attribute_code"
}

func (store AttrCodeS) TransformToView() AttrCodeV {
	return AttrCodeV{
		ID:    store.ID,
		Title: store.Title,
	}
}

type AttrCodeV struct {
	ID    int64
	Title string
}

func (view AttrCodeV) TransformToStore() AttrCodeS {
	return AttrCodeS{
		ID:    view.ID,
		Title: view.Title,
	}
}
