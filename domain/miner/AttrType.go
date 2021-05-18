package miner

type AttrTypeS struct {
	ID    int64
	Title string
}

func (AttrTypeS) TableName() string {
	return "attribute_type"
}

func (store AttrTypeS) TransformToView() AttrTypeV {
	return AttrTypeV{
		ID:    store.ID,
		Title: store.Title,
	}
}

type AttrTypeV struct {
	ID    int64
	Title string
}

func (view AttrTypeV) TransformToStore() AttrTypeS {
	return AttrTypeS{
		ID:    view.ID,
		Title: view.Title,
	}
}
