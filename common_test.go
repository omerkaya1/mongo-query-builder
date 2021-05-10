package mongo_query_builder

import (
	"reflect"
	"testing"
)

func TestComposeNestedField(t *testing.T) {
	for _, tc := range []struct {
		header string
		input  []string
		output string
	}{
		{"Empty", nil, ""},
		{"Single", []string{"one"}, "one"},
		{"Multiple", []string{"one", "two", "three"}, "one.two.three"},
	} {
		t.Run(tc.header, func(t *testing.T) {
			reflect.DeepEqual(tc.output, ComposeNestedField(tc.input...))
		})
	}
}
