package utils

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// Decode is a wrapper around mapstructure.Decode
func Decode(data any, response any) error {
	return mapstructure.WeakDecode(data, response)
}

// DecodeMap is a wrapper around mapstructure.WeakDecode for multiple
// values passed as a map[string]interface{}
func DecodeMap[V any](data map[string]any, response map[int]V) error {
	for k, v := range data {
		var d V
		err := Decode(v, &d)
		if err != nil {
			return err
		}

		key, err := strconv.Atoi(k)
		if err != nil {
			return err
		}

		response[key] = d
	}

	return nil
}

// DecodeSlice is a wrapper around mapstructure.WeakDecode for multiple
// values passed as a []interface{} - response must be a pointer to a slices
func DecodeSlice[V any](data []any, response *[]V) error {
	for _, v := range data {
		var d V
		err := Decode(v, &d)
		if err != nil {
			return err
		}

		*response = append(*response, d)
	}

	return nil
}
