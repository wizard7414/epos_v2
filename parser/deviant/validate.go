package deviant

import "github.com/wizard7414/epos_v2/domain/miner"

func validateSourceResource(resource miner.ResourceV) bool {
	return resource.Title != "" && resource.Url != ""
}

func validateEntrySource(source miner.SourceV) bool {
	return source.Title != "" && source.Url != "" && validateSourceResource(source.Resource)
}

func validateObjectEntry(entry miner.EntryV) bool {
	return entry.Title != "" && entry.Url != "" && validateEntrySource(entry.Source)
}

func validateObjectType(objectType miner.ObjTypeV) bool {
	return objectType.Title != ""
}

func validateObjectAttributes(attributes []miner.AttrV) bool {
	if len(attributes) == 0 {
		return false
	}
	result := true
	for id := range attributes {
		attribute := attributes[id]
		result = result && attribute.Value != "" && attribute.Code.Title != "" && attribute.AttrType.Title != ""
	}
	return result
}

func validateObjectMeta(object miner.ObjectV) bool {
	return validateObjectEntry(object.Entry) && validateObjectType(object.ObjectType) && validateObjectAttributes(object.Attributes)
}

func ValidateParsedObject(object miner.ObjectV) bool {
	return object.Title != "" && object.Url != "" && validateObjectMeta(object)
}
