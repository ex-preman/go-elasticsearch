package types

import (
	jsoniter "github.com/json-iterator/go"
)

type Decoder struct{}

func (u *Decoder) Decode(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}

func (u *Decoder) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}

func (u *Decoder) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(v)
}
