package mongo_query_builder

import "gopkg.in/mgo.v2/bson"

// MgoQueryBuilder is used for composing the DB queries for the community-supported mgo package
type MgoQueryBuilder struct {
	query     []bson.DocElem
	selection []bson.DocElem
	sort      []bson.DocElem
	limit     int64
}

// NewMgoQueryBuilder returns new MgoQueryBuilderBuilder object to its caller
func NewMgoQueryBuilder() *MgoQueryBuilder {
	return &MgoQueryBuilder{}
}

// Clear restores the initial state of the query
func (mqb *MgoQueryBuilder) Clear() {
	mqb.query = mqb.query[:0]
	mqb.selection = mqb.selection[:0]
	mqb.sort = mqb.sort[:0]
}

func (mqb *MgoQueryBuilder) Query() interface{} {
	return mqb.query
}

func (mqb *MgoQueryBuilder) Selection() interface{} {
	return mqb.selection
}

func (mqb *MgoQueryBuilder) Limit() int64 {
	return mqb.limit
}

func (mqb *MgoQueryBuilder) Sort() interface{} {
	return mqb.sort
}

func (mqb *MgoQueryBuilder) AddQueryItem(field string, value string) *MgoQueryBuilder {
	mqb.query = append(mqb.query, bson.DocElem{Name: field, Value: value})
	return mqb
}

func (mqb *MgoQueryBuilder) AddInternalID(internalID string) *MgoQueryBuilder {
	if bson.IsObjectIdHex(internalID) {
		mqb.query = append(mqb.query, bson.DocElem{Name: idField, Value: bson.ObjectIdHex(internalID)})
	}
	return mqb
}

func (mqb *MgoQueryBuilder) AddInternalObjectID(id bson.ObjectId) *MgoQueryBuilder {
	mqb.query = append(mqb.query, bson.DocElem{Name: idField, Value: id})
	return mqb
}

func (mqb *MgoQueryBuilder) AddLimit(limit int64) *MgoQueryBuilder {
	mqb.limit = limit
	return mqb
}

func (mqb *MgoQueryBuilder) AddOrder(orderField string, asc bool) *MgoQueryBuilder {
	if !asc {
		mqb.sort = append(mqb.sort, bson.DocElem{Name: orderField, Value: -1})
	} else {
		mqb.sort = append(mqb.sort, bson.DocElem{Name: orderField, Value: 1})
	}
	return mqb
}

func (mqb *MgoQueryBuilder) AddSelectionItem(fields ...string) *MgoQueryBuilder {
	for i := range fields {
		mqb.selection = append(mqb.selection, bson.DocElem{Name: fields[i], Value: 1})
	}
	return mqb
}

func (mqb *MgoQueryBuilder) AddDeselectionItem(fields ...string) *MgoQueryBuilder {
	for i := range fields {
		mqb.selection = append(mqb.selection, bson.DocElem{Name: fields[i], Value: -1})
	}
	return mqb
}

func (mqb *MgoQueryBuilder) AddRegexp(field, regexStr string) *MgoQueryBuilder {
	mqb.query = append(mqb.query, bson.DocElem{Name: field, Value: bson.DocElem{Name: regex, Value: regexStr}})
	return mqb
}
