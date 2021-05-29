package dbMiner

import "github.com/wizard7414/epos_v2/domain/miner"

type EposDb interface {
	InitDb(path string) error

	CloseDb() error

	CreateAttribute(attribute miner.AttrV) error

	CreateSource(source miner.SourceV) error

	CreateEntry(entry miner.EntryV) error

	CreateObject(object miner.ObjectV) error

	GetObjectAttributes(objectId int64) ([]miner.AttrV, error)

	GetObjectByUrl(objectUrl string) ([]miner.ObjectV, error)
}

type AttributeCodeDao interface {
	Create(attributeCode miner.AttrCodeS) error

	Delete(attributeCodeId int64) error

	GetById(attributeCodeId int64) (miner.AttrCodeS, error)

	GetByTitle(attributeCodeTitle string) ([]miner.AttrCodeS, error)
}

type AttributeTypeDao interface {
	Create(attributeType miner.AttrTypeS) error

	Delete(attributeTypeId int64) error

	GetById(attributeTypeId int64) (miner.AttrTypeS, error)

	GetByTitle(attributeTypeTitle string) ([]miner.AttrTypeS, error)
}

type AttributeDao interface {
	Create(attribute miner.AttrS) error

	Delete(attributeId int64) error

	GetById(attributeId int64) (miner.AttrS, error)

	GetByCode(attributeCode int64) ([]miner.AttrS, error)
}

type ResourceDao interface {
	Create(resource miner.ResourceS) error

	Delete(resourceId int64) error

	GetById(resourceId int64) (miner.ResourceS, error)

	GetByTitle(resourceTitle string) ([]miner.ResourceS, error)
}

type SourceDao interface {
	Create(source miner.SourceS) error

	Delete(sourceId int64) error

	GetById(sourceId int64) (miner.SourceS, error)

	GetByTitle(sourceTitle string) ([]miner.SourceS, error)

	UpdateDateById(sourceId int64, newDate int64) error

	UpdateDateByTitle(sourceTitle string, newDate int64) error
}

type EntryDao interface {
	Create(entry miner.EntryS) error

	Delete(entryId int64) error

	GetById(entryId int64) (miner.EntryS, error)

	GetByTitle(entryTitle string) ([]miner.EntryS, error)
}

type ObjectTypeDao interface {
	Create(objectType miner.ObjTypeS) error

	Delete(objectTypeId int64) error

	GetById(objectTypeId int64) (miner.ObjTypeS, error)

	GetByTitle(objectTypeTitle string) ([]miner.ObjTypeS, error)
}

type ObjectDao interface {
	Create(object miner.ObjectS) error

	Delete(objectId int64) error

	GetById(objectId int64) (miner.ObjectS, error)

	GetByTitle(objectTitle string) ([]miner.ObjectS, error)

	GetByUrl(objectUrl string) ([]miner.ObjectS, error)
}

type ObjectAttributeDao interface {
	Create(objectAttribute miner.ObjAttrS) error

	Delete(objectId int64, attributeId int64) error

	GetById(objectId int64, attributeId int64) (miner.ObjAttrS, error)

	GetByObject(objectId int64) ([]miner.ObjAttrS, error)

	GetByAttribute(attributeId int64) ([]miner.ObjAttrS, error)
}
