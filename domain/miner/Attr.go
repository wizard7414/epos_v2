package miner

type AttrS struct {
	ID            int64
	Code          int64
	AttributeType int64
	Value         string
}

func (AttrS) TableName() string {
	return "attribute"
}

func (store AttrS) TransformToView() AttrV {
	return AttrV{
		ID: store.ID,
		Code: AttrCodeV{
			ID:    store.Code,
			Title: "",
		},
		AttrType: AttrTypeV{
			ID:    store.AttributeType,
			Title: "",
		},
		Value: store.Value,
	}
}

type AttrV struct {
	ID       int64
	Code     AttrCodeV
	AttrType AttrTypeV
	Value    string
}

func (view AttrV) TransformToStore() AttrS {
	return AttrS{
		ID:            view.ID,
		Code:          view.Code.ID,
		AttributeType: view.AttrType.ID,
		Value:         view.Value,
	}
}
