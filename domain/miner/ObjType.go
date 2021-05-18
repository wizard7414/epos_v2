package miner

type ObjTypeS struct {
	ID    int64
	Title string
}

func (ObjTypeS) TableName() string {
	return "object_type"
}

func (store ObjTypeS) TransformToView() ObjTypeV {
	return ObjTypeV{
		ID:    store.ID,
		Title: store.Title,
	}
}

type ObjTypeV struct {
	ID    int64
	Title string
}

func (view ObjTypeV) TransformToStore() ObjTypeS {
	return ObjTypeS{
		ID:    view.ID,
		Title: view.Title,
	}
}
