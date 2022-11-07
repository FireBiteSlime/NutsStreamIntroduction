package deserializer

import (
	"encoding/json"
	"nstream/models"
)

func MarshalJson(order models.Order) ([]byte, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return data, err
	}
	return data, nil
}
