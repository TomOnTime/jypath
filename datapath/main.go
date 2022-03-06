package datapath

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Item struct {
	Key string
	Val interface{}
}

func FromJSON(data []byte) (interface{}, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	return v, err
}

func FromPaths(b []byte) interface{} { return nil }

func FromYAML(data []byte) (interface{}, error) {
	var v interface{}
	err := yaml.Unmarshal(data, &v)
	return v, err
}

func ToJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func ToPaths(i interface{}) []Item       { return nil }
func ToPathsString(i interface{}) string { return "" }
func ToYAML(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
