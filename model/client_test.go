package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClientUpdateJsonMarshal(t *testing.T) {
	var cu ClientUpdate
	emptyJSON(t, cu)
}

func emptyJSON(t *testing.T, val interface{}) {
	var emptyJSON = []byte("{}")
	data, _ := json.MarshalIndent(val, "", " ")
	if !reflect.DeepEqual(data, emptyJSON) {
		t.Errorf("for empty rendering JSON should find {} got %v", string(data))
	}
}
