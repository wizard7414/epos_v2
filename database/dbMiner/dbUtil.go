package dbMiner

import (
	miner2 "github.com/wizard7414/epos_v2/domain/miner"
	"time"
)

//func Init(configuration domain.EposConfig) (*database.EposDatabase, error) {
//	database := database.EposDatabase{
//		DbMiner: DbMiner{
//			AttributeCode:   AttrCodeDao{
//				DB: nil,
//			},
//			AttributeType:   AttrTypeDao{
//				DB: nil,
//			},
//			Attribute:       AttrDao{
//				DB: nil,
//			},
//			Resource:        ResDao{
//				DB: nil,
//			},
//			Source:          SrcDao{
//				DB: nil,
//			},
//			Entry:           EntDao{
//				DB: nil,
//			},
//			ObjectType:      ObjTypeDao{
//				DB: nil,
//			},
//			Object:          ObjDao{
//				DB: nil,
//			},
//			ObjectAttribute: ObjAttrDao{
//				DB: nil,
//			},
//		},
//	}
//
//	err := database.DbMiner.InitDb(configuration.DbMinerPath)
//	return &database, err
//}

func CreateOrGetAttrCode(db *DbMiner, attributeCode miner2.AttrCodeV) (miner2.AttrCodeS, error) {
	attributeCodes, findErr := db.AttributeCode.GetByTitle(attributeCode.Title)
	var updatedAttributeCode miner2.AttrCodeS

	if findErr != nil {
		return updatedAttributeCode, findErr
	} else if len(attributeCodes) == 0 {
		attributeCode.ID = time.Now().UnixNano()
		createErr := db.AttributeCode.Create(attributeCode.TransformToStore())
		if createErr != nil {
			return updatedAttributeCode, createErr
		} else {
			updatedAttributeCode = attributeCode.TransformToStore()
		}
	} else if len(attributeCodes) == 1 {
		updatedAttributeCode = attributeCodes[0]
	}
	return updatedAttributeCode, nil
}

func CreateOrGetAttrType(db *DbMiner, attributeType miner2.AttrTypeV) (miner2.AttrTypeS, error) {
	attributeTypes, findErr := db.AttributeType.GetByTitle(attributeType.Title)
	var updatedAttributeType miner2.AttrTypeS

	if findErr != nil {
		return updatedAttributeType, findErr
	} else if len(attributeTypes) == 0 {
		attributeType.ID = time.Now().UnixNano()
		createErr := db.AttributeType.Create(attributeType.TransformToStore())
		if createErr != nil {
			return updatedAttributeType, createErr
		} else {
			updatedAttributeType = attributeType.TransformToStore()
		}
	} else if len(attributeTypes) == 1 {
		updatedAttributeType = attributeTypes[0]
	}
	return updatedAttributeType, nil
}

func CreateOrGetRes(db *DbMiner, resource miner2.ResourceV) (miner2.ResourceS, error) {
	resources, findErr := db.Resource.GetByTitle(resource.Title)
	var updatedResource miner2.ResourceS

	if findErr != nil {
		return updatedResource, findErr
	} else if len(resources) == 0 {
		resource.ID = time.Now().UnixNano()
		createErr := db.Resource.Create(resource.TransformToStore())
		if createErr != nil {
			return updatedResource, createErr
		} else {
			updatedResource = resource.TransformToStore()
		}
	} else if len(resources) == 1 {
		updatedResource = resources[0]
	}
	return updatedResource, nil
}

func CreateOrGetSource(db *DbMiner, source miner2.SourceV) (miner2.SourceS, error) {
	sources, findErr := db.Source.GetByTitle(source.Title)
	var updatedSource miner2.SourceS

	if findErr != nil {
		return updatedSource, findErr
	} else if len(sources) == 0 {
		source.ID = time.Now().UnixNano()
		createErr := db.CreateSource(source)
		if createErr != nil {
			return updatedSource, createErr
		} else {
			updatedSource = source.TransformToStore()
		}
	} else if len(sources) == 1 {
		updatedSource = sources[0]
	}
	return updatedSource, nil
}

func CreateOrGetObjType(db *DbMiner, objectType miner2.ObjTypeV) (miner2.ObjTypeS, error) {
	objectTypes, findErr := db.ObjectType.GetByTitle(objectType.Title)
	var updateObjectType miner2.ObjTypeS

	if findErr != nil {
		return updateObjectType, findErr
	} else if len(objectTypes) == 0 {
		objectType.ID = time.Now().UnixNano()
		createErr := db.ObjectType.Create(objectType.TransformToStore())
		if createErr != nil {
			return updateObjectType, createErr
		} else {
			updateObjectType = objectType.TransformToStore()
		}
	} else if len(objectTypes) == 1 {
		updateObjectType = objectTypes[0]
	}
	return updateObjectType, nil
}
