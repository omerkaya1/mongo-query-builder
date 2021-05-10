package mongo_query_builder

import "strings"

const (
	regex   = "$regex"
	idField = "_id"
)

// ComposeNestedField accepts a slice of strings which represent a nested path to a field in a document for querying and
// returns a properly encoded string version of it
func ComposeNestedField(fields ...string) string {
	return strings.Join(fields, ".")
}
