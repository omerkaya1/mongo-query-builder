package mongo_query_builder

import "go.mongodb.org/mongo-driver/bson/primitive"

// MongoQueryBuilder is used for composing the DB queries for the standard mongo-driver
type MongoQueryBuilder struct {
	query     primitive.D
	selection primitive.D
	sort      primitive.D
	limit     int64
}

// NewMongoQueryBuilder returns a new MongoQueryBuilder object to its caller
func NewMongoQueryBuilder() *MongoQueryBuilder {
	return &MongoQueryBuilder{}
}

// Clear restores the initial state of the query
func (mqb *MongoQueryBuilder) Clear() {
	mqb.query = mqb.query[:0]
	mqb.selection = mqb.selection[:0]
	mqb.sort = mqb.sort[:0]
}

func (mqb *MongoQueryBuilder) AddSelectionItem(fields ...string) *MongoQueryBuilder {
	for i := range fields {
		mqb.selection = append(mqb.selection, primitive.E{Key: fields[i], Value: 1})
	}
	return mqb
}

func (mqb *MongoQueryBuilder) AddDeselectionItem(fields ...string) *MongoQueryBuilder {
	for i := range fields {
		mqb.selection = append(mqb.selection, primitive.E{Key: fields[i], Value: 0})
	}
	return mqb
}

func (mqb *MongoQueryBuilder) Query() interface{} {
	return mqb.query
}

func (mqb *MongoQueryBuilder) Selection() interface{} {
	return mqb.selection
}

func (mqb *MongoQueryBuilder) Limit() int64 {
	return mqb.limit
}

func (mqb *MongoQueryBuilder) Sort() interface{} {
	return mqb.sort
}

func (mqb *MongoQueryBuilder) AddQueryItem(field string, value interface{}) *MongoQueryBuilder {
	mqb.query = append(mqb.query, primitive.E{Key: field, Value: value})
	return mqb
}

func (mqb *MongoQueryBuilder) AddInternalID(internalID string) *MongoQueryBuilder {
	if id, err := primitive.ObjectIDFromHex(internalID); err == nil {
		mqb.query = append(mqb.query, primitive.E{Key: idField, Value: id})
	}
	return mqb
}

func (mqb *MongoQueryBuilder) AddInternalObjectID(id primitive.ObjectID) *MongoQueryBuilder {
	mqb.query = append(mqb.query, primitive.E{Key: idField, Value: id})
	return mqb
}

func (mqb *MongoQueryBuilder) AddLimit(limit int64) *MongoQueryBuilder {
	mqb.limit = limit
	return mqb
}

func (mqb *MongoQueryBuilder) AddOrder(field string, asc bool) *MongoQueryBuilder {
	if !asc {
		mqb.sort = append(mqb.sort, primitive.E{Key: field, Value: -1})
	} else {
		mqb.sort = append(mqb.sort, primitive.E{Key: field, Value: 1})
	}
	return mqb
}

func (mqb *MongoQueryBuilder) AddRegexp(field, regexStr string) *MongoQueryBuilder {
	mqb.query = append(mqb.query, primitive.E{Key: field, Value: primitive.E{Key: regex, Value: regexStr}})
	return mqb
}
